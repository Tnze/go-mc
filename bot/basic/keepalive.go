package basic

import (
	"time"

	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

const keepAliveDuration = time.Second * 20

func (p *Player) resetKeepAliveDeadline() {
	newDeadline := time.Now().Add(keepAliveDuration)
	p.c.Conn.Socket.SetDeadline(newDeadline)
}

func (p *Player) handleKeepAlivePacket(packet pk.Packet) error {
	var KeepAliveID pk.Long
	if err := packet.Scan(&KeepAliveID); err != nil {
		return Error{err}
	}

	p.resetKeepAliveDeadline()

	// Response
	err := p.c.Conn.WritePacket(pk.Packet{
		ID:   int32(packetid.ServerboundKeepAlive),
		Data: packet.Data,
	})
	if err != nil {
		return Error{err}
	}
	return nil
}
