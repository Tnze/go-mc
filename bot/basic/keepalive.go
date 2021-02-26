package basic

import (
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

func (p Player) handleKeepAlivePacket(packet pk.Packet) error {
	var KeepAliveID pk.Long
	if err := packet.Scan(&KeepAliveID); err != nil {
		return err
	}
	// Response
	return p.c.Conn.WritePacket(pk.Marshal(
		packetid.KeepAliveServerbound,
		KeepAliveID,
	))
}
