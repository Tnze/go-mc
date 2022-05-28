package command

import (
	"context"
	"strings"

	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/server"
)

// Init implement server.Component for Graph
func (g *Graph) Init(game *server.Game) {
	game.AddHandler(&server.PacketHandler{
		ID: packetid.ServerboundChat,
		F: func(client *server.Client, player *server.Player, packet server.Packet758) error {
			var msg pk.String
			if err := pk.Packet(packet).Scan(&msg); err != nil {
				return err
			}
			if cmd := string(msg); strings.HasPrefix(cmd, "/") {
				ctx := context.WithValue(context.TODO(), "sender", player)
				cmderr := g.Execute(ctx, strings.TrimPrefix(cmd, "/"))
				if cmderr != nil {
					// TODO: tell player that their command has error
				}
			}
			return nil
		},
	})
}

// Run implement server.Component for Graph
func (g *Graph) Run(ctx context.Context) {}

// ClientJoin implement server.Component for Graph
func (g *Graph) ClientJoin(client *server.Client, _ *server.Player) {
	client.WritePacket(server.Packet758(pk.Marshal(
		packetid.ClientboundCommands, g,
	)))
}

// ClientLeft implement server.Component for Graph
func (g *Graph) ClientLeft(_ *server.Client, _ *server.Player, _ error) {}
