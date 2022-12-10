package command

import (
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Client interface {
	SendPacket(p pk.Packet)
}

// ClientJoin implement server.Component for Graph
func (g *Graph) ClientJoin(client Client) {
	client.SendPacket(pk.Marshal(
		packetid.CPacketCommands, g,
	))
}
