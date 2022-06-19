package server

import (
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"time"
	"unsafe"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/offline"
	"github.com/Tnze/go-mc/server/auth"
)

// LoginHandler is used to handle player login process, that is,
// from clientbound "LoginStart" packet to serverbound "LoginSuccess" packet.
type LoginHandler interface {
	AcceptLogin(conn *net.Conn, protocol int32) (name string, id uuid.UUID, profilePubKey *rsa.PublicKey, err error)
}

// LoginChecker is the interface to check if a player is allowed to log in the server.
// The checking could be anything, server player number, protocol version, blacklist or whitelist.
// If a player is not allowed to, the reason should be returned and will be sent to client by "LoginDisconnect" packet.
type LoginChecker interface {
	CheckPlayer(name string, id uuid.UUID, protocol int32) (ok bool, reason chat.Message)
}

// MojangLoginHandler is a standard LoginHandler that implement both online and offline login progress.
// This implementation also support custom LoginChecker.
// None of Custom login packet (also called LoginPluginRequest/Response) is support by this implementation.
// To do that, implement your own LoginHandler imitate this code.
type MojangLoginHandler struct {
	// OnlineMode enables to check player's account.
	// And also encrypt the connection after login.
	OnlineMode bool

	// EnforceSecureProfile enforce to check the player's profile public key
	EnforceSecureProfile bool

	// Threshold set the smallest size of raw network payload to compress.
	// Set to 0 to compress all packets. Set to -1 to disable compression.
	Threshold int

	// LoginChecker is used to apply some checks before sending "LoginSuccess" packet
	// (e.g. blacklist or is server full).
	// This is optional field and can be set to nil.
	LoginChecker
}

// AcceptLogin implement LoginHandler for MojangLoginHandler
func (d *MojangLoginHandler) AcceptLogin(conn *net.Conn, protocol int32) (name string, id uuid.UUID, profilePubKey *rsa.PublicKey, err error) {
	//login start
	var p pk.Packet
	err = conn.ReadPacket(&p)
	if err != nil {
		return
	}
	if p.ID != packetid.LoginStart {
		err = wrongPacketErr{expect: packetid.LoginStart, get: p.ID}
		return
	}

	var hasSignature pk.Boolean
	var timestamp pk.Long
	var profilePubKeyBytes pk.ByteArray
	var signature pk.ByteArray
	err = p.Scan(
		(*pk.String)(&name),
		&hasSignature, pk.Opt{
			Has:   &hasSignature,
			Field: pk.Tuple{&timestamp, &profilePubKeyBytes, &signature},
		},
	) //decode username as pk.String
	if err != nil {
		return
	}

	if hasSignature {
		if time.UnixMilli(int64(timestamp)).Before(time.Now()) ||
			!auth.VerifySignature(profilePubKeyBytes, signature) {
			err = LoginFailErr{reason: chat.TranslateMsg("multiplayer.disconnect.invalid_public_key_signature")}
			return
		}
	} else if d.EnforceSecureProfile {
		err = LoginFailErr{reason: chat.TranslateMsg("multiplayer.disconnect.missing_public_key")}
		return
	}

	var properties []property
	//auth
	if d.OnlineMode {
		var resp *auth.Resp
		//Auth, Encrypt
		var key any
		key, err = x509.ParsePKIXPublicKey(profilePubKeyBytes)
		if err != nil {
			return
		}
		if rsaPubKey, ok := key.(*rsa.PublicKey); !ok {
			err = errors.New("expect RSA public key")
			return
		} else {
			profilePubKey = rsaPubKey
		}

		resp, err = auth.Encrypt(conn, name, profilePubKey)
		if err != nil {
			return
		}
		name = resp.Name
		id = resp.ID
		properties = *(*[]property)(unsafe.Pointer(&resp.Properties))
	} else {
		// offline-mode UUID
		id = offline.NameToUUID(name)
	}

	//set compression
	if d.Threshold >= 0 {
		err = conn.WritePacket(pk.Marshal(
			packetid.LoginCompression,
			pk.VarInt(d.Threshold),
		))
		if err != nil {
			return
		}
		conn.SetThreshold(d.Threshold)
	}

	// check if player can join (whitelist, blacklist, server full or something else)
	if d.LoginChecker != nil {
		if ok, result := d.CheckPlayer(name, id, protocol); !ok {
			// player is not allowed to join the server
			err = LoginFailErr{reason: result}
			return
		}
	}
	// send login success
	err = conn.WritePacket(pk.Marshal(
		packetid.LoginSuccess,
		pk.UUID(id),
		pk.String(name),
		pk.Array(properties),
	))
	return
}

type GameProfile struct {
	ID   uuid.UUID
	Name string
}

type property auth.Property

func (p property) WriteTo(w io.Writer) (n int64, err error) {
	hasSignature := len(p.Signature) > 0
	return pk.Tuple{
		pk.String(p.Name),
		pk.String(p.Value),
		pk.Boolean(hasSignature),
		pk.Opt{
			Has:   hasSignature,
			Field: pk.String(p.Signature),
		},
	}.WriteTo(w)
}

type wrongPacketErr struct {
	expect, get int32
}

func (w wrongPacketErr) Error() string {
	return fmt.Sprintf("wrong packet id: expect %#02X, get %#02X", w.expect, w.get)
}

type LoginFailErr struct {
	reason chat.Message
}

func (l LoginFailErr) Error() string {
	return "login error: " + l.reason.ClearString()
}
