package server

import (
	"fmt"

	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/offline"
	"github.com/Tnze/go-mc/server/auth"
	"github.com/google/uuid"
)

type LoginHandler interface {
	AcceptLogin(conn *net.Conn, protocol int32) (name string, id uuid.UUID, err error)
}

type MojangLoginHandler struct {
	OnlineMode bool
	Threshold  int
}

func (d *MojangLoginHandler) AcceptLogin(conn *net.Conn, protocol int32) (name string, id uuid.UUID, err error) {
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

	err = p.Scan((*pk.String)(&name)) //decode username as pk.String
	if err != nil {
		return
	}

	//auth
	if d.OnlineMode {
		var resp *auth.Resp
		//Auth, Encrypt
		resp, err = auth.Encrypt(conn, name)
		if err != nil {
			return
		}
		name = resp.Name
		id = resp.ID
	} else {
		// offline-mode UUID
		id = offline.NameToUUID(name)
	}

	//set compression
	if d.Threshold >= 0 {
		err = conn.WritePacket(pk.Marshal(
			packetid.SetCompression, pk.VarInt(d.Threshold),
		))
		if err != nil {
			return
		}
		conn.SetThreshold(d.Threshold)
	}

	// send login success
	err = conn.WritePacket(pk.Marshal(packetid.LoginSuccess,
		pk.UUID(id),
		pk.String(name),
	))
	return
}

type wrongPacketErr struct {
	expect, get int32
}

func (w wrongPacketErr) Error() string {
	return fmt.Sprintf("wrong packet id: expect %#02X, get %#02X", w.expect, w.get)
}
