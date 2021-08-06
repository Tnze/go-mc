package bot

import (
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

func (c Client) handleKeepAlivePacket(packet pk.Packet) error {
	var KeepAliveID pk.Long
	if err := packet.Scan(&KeepAliveID); err != nil {
		return err
	}
	// Response
	err := c.Conn.WritePacket(pk.Packet{
		ID:   packetid.KeepAliveServerbound,
		Data: packet.Data,
	})
	if err != nil {
		return err
	}
	return nil
}
