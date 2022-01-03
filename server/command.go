package server

import (
	"context"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/server/command"
	"strings"
)

type CommandGraph struct {
	*command.Graph
}

func NewCommandGraph() *CommandGraph {
	return &CommandGraph{
		Graph: command.NewGraph(),
	}
}

func (c *CommandGraph) Init(g *Game) {
	g.AddHandler(&PacketHandler{
		ID: packetid.ServerboundChat,
		F: func(player *Player, packet Packet757) error {
			var msg pk.String
			if err := pk.Packet(packet).Scan(&msg); err != nil {
				return err
			}
			if cmd := string(msg); strings.HasPrefix(cmd, "/") {
				cmderr := c.Graph.Run(strings.TrimPrefix(cmd, "/"))
				if cmderr != nil {
					// TODO: tell player that their command has error
				}
			}
			return nil
		},
	})
}

func (c *CommandGraph) Run(ctx context.Context) {}

func (c *CommandGraph) AddPlayer(p *Player) {
	p.WritePacket(Packet757(pk.Marshal(
		packetid.ClientboundCommands,
		c.Graph,
	)))
}

func (c *CommandGraph) RemovePlayer(p *Player) {}
