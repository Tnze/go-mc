package server

import (
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Player struct {
	*net.Conn
	EntityID int32
	Gamemode byte
}

// Packet757 is a packet in protocol 757.
// We are using type system to force programmers to update packets.
type Packet757 pk.Packet

// WritePacket to player client. The type of parameter will update per version.
func (p *Player) WritePacket(packet Packet757) error {
	return p.Conn.WritePacket(pk.Packet(packet))
}
