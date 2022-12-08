// Package bot implements a simple Minecraft client that can join a server
// or just ping it for getting information.
//
// Runnable example could be found at examples/ .
package bot

import (
	"context"
	"errors"
	"net"
	"strconv"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	mcnet "github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/yggdrasil/user"
	"github.com/google/uuid"
)

// ProtocolVersion is the protocol version number of minecraft net protocol
const (
	ProtocolVersion = 761
	DefaultPort     = mcnet.DefaultPort
)

type JoinOptions struct {
	Dialer  *net.Dialer
	Context context.Context

	// Indicate not to fetch and sending player's PubKey
	NoPublicKey bool

	// Specify the player PubKey to use.
	// If nil, it will be obtained from Mojang when joining
	KeyPair *user.KeyPairResp
}

// JoinServer connect a Minecraft server for playing the game.
// Using roughly the same way to parse address as minecraft.
func (c *Client) JoinServer(addr string) (err error) {
	return c.join(addr, JoinOptions{
		Context: context.Background(),
		Dialer:  (*net.Dialer)(&mcnet.DefaultDialer),
	})
}

// JoinServerWithDialer is similar to JoinServer but using a Dialer.
func (c *Client) JoinServerWithDialer(dialer *net.Dialer, addr string) (err error) {
	return c.join(addr, JoinOptions{
		Context: context.Background(),
		Dialer:  dialer,
	})
}

func (c *Client) JoinServerWithOptions(addr string, options JoinOptions) (err error) {
	if options.Dialer == nil {
		options.Dialer = (*net.Dialer)(&mcnet.DefaultDialer)
	}
	if options.Context == nil {
		options.Context = context.Background()
	}
	return c.join(addr, options)
}

func (c *Client) join(addr string, options JoinOptions) error {
	const Handshake = 0x00
	// Split Host and Port
	host, portStr, err := net.SplitHostPort(addr)
	var port uint64
	if err != nil {
		var addrErr *net.AddrError
		const missingPort = "missing port in address"
		if errors.As(err, &addrErr) && addrErr.Err == missingPort {
			host = addr
			port = 25565
		} else {
			return LoginErr{"split address", err}
		}
	} else {
		port, err = strconv.ParseUint(portStr, 0, 16)
		if err != nil {
			return LoginErr{"parse port", err}
		}
	}

	// Dial connection
	d := (*mcnet.Dialer)(options.Dialer)
	ctx := options.Context
	c.Conn, err = d.DialMCContext(ctx, addr)
	if err != nil {
		return LoginErr{"connect server", err}
	}

	// Handshake
	err = c.Conn.WritePacket(pk.Marshal(
		Handshake,
		pk.VarInt(ProtocolVersion), // Protocol version
		pk.String(host),            // Host
		pk.UnsignedShort(port),     // Port
		pk.VarInt(2),
	))
	if err != nil {
		return LoginErr{"handshake", err}
	}
	// Login Start
	if c.Auth.AsTk != "" && !options.NoPublicKey {
		if options.KeyPair != nil {
			c.KeyPair = options.KeyPair
		} else if KeyPairResp, err := user.GetOrFetchKeyPair(c.Auth.AsTk); err == nil {
			c.KeyPair = &KeyPairResp
		}
	}
	c.UUID, err = uuid.Parse(c.Auth.UUID)
	PlayerUUID := pk.Option[pk.UUID, *pk.UUID]{
		Has: err == nil,
		Val: pk.UUID(c.UUID),
	}
	err = c.Conn.WritePacket(pk.Marshal(
		packetid.LoginStart,
		pk.String(c.Auth.Name),
		PlayerUUID,
	))
	if err != nil {
		return LoginErr{"login start", err}
	}
	for {
		// Receive Packet
		var p pk.Packet
		if err = c.Conn.ReadPacket(&p); err != nil {
			return LoginErr{"receive packet", err}
		}

		// Handle Packet
		switch p.ID {
		case packetid.LoginDisconnect: // LoginDisconnect
			var reason chat.Message
			err = p.Scan(&reason)
			if err != nil {
				return LoginErr{"disconnect", err}
			}
			return LoginErr{"disconnect", DisconnectErr(reason)}

		case packetid.LoginEncryptionRequest: // Encryption Request
			if err := handleEncryptionRequest(c, p); err != nil {
				return LoginErr{"encryption", err}
			}

		case packetid.LoginSuccess: // Login Success
			err := p.Scan(
				(*pk.UUID)(&c.UUID),
				(*pk.String)(&c.Name),
			)
			if err != nil {
				return LoginErr{"login success", err}
			}
			return nil

		case packetid.LoginCompression: // Set Compression
			var threshold pk.VarInt
			if err := p.Scan(&threshold); err != nil {
				return LoginErr{"compression", err}
			}
			c.Conn.SetThreshold(int(threshold))

		case packetid.LoginPluginRequest: // Login Plugin Request
			var (
				msgid   pk.VarInt
				channel pk.Identifier
				data    pk.PluginMessageData
			)
			if err := p.Scan(&msgid, &channel, &data); err != nil {
				return LoginErr{"Login Plugin", err}
			}

			var PluginMessageData pk.Option[pk.PluginMessageData, *pk.PluginMessageData]
			if handler, ok := c.LoginPlugin[string(channel)]; ok {
				PluginMessageData.Has = true
				PluginMessageData.Val, err = handler(data)
				if err != nil {
					return LoginErr{"Login Plugin", err}
				}
			}

			if err := c.Conn.WritePacket(pk.Marshal(
				packetid.LoginPluginResponse,
				msgid, PluginMessageData,
			)); err != nil {
				return LoginErr{"login Plugin", err}
			}
		}
	}
}

type LoginErr struct {
	Stage string
	Err   error
}

func (l LoginErr) Error() string {
	return "bot: " + l.Stage + " error: " + l.Err.Error()
}

func (l LoginErr) Unwrap() error {
	return l.Err
}

type DisconnectErr chat.Message

func (d DisconnectErr) Error() string {
	return "disconnect because: " + chat.Message(d).String()
}
