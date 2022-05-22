package clientinfo

import (
	"context"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/server"
)

type ClientInformation struct {
	Players map[uuid.UUID]*Info
}

type Info struct {
	Locale              string
	ViewDistance        int
	ChatMode            byte
	ChatColors          bool
	DisplayedSkinParts  byte
	MainHand            byte // 0: Left, 1: Right.
	EnableTextFiltering bool
	AllowServerListings bool
}

func (c *ClientInformation) Init(g *server.Game) {
	c.Players = make(map[uuid.UUID]*Info)
	g.AddHandler(&server.PacketHandler{
		ID: packetid.ServerboundClientInformation,
		F: func(player *server.Player, p server.Packet758) error {
			var (
				Locale              pk.String
				ViewDistance        pk.Byte
				ChatMode            pk.VarInt
				ChatColors          pk.Boolean
				DisplayedSkinParts  pk.UnsignedByte
				MainHand            pk.VarInt
				EnableTextFiltering pk.Boolean
				AllowServerListings pk.Boolean
			)
			err := pk.Packet(p).Scan(
				&Locale,
				&ViewDistance,
				&ChatMode,
				&ChatColors,
				&DisplayedSkinParts,
				&MainHand,
				&EnableTextFiltering,
				&AllowServerListings,
			)
			if err != nil {
				return err
			}
			c.Players[player.UUID] = &Info{
				Locale:              string(Locale),
				ViewDistance:        int(ViewDistance),
				ChatMode:            byte(ChatMode),
				ChatColors:          bool(ChatColors),
				DisplayedSkinParts:  byte(DisplayedSkinParts),
				MainHand:            byte(MainHand),
				EnableTextFiltering: bool(EnableTextFiltering),
				AllowServerListings: bool(AllowServerListings),
			}
			return nil
		},
	})
}

func (c *ClientInformation) Run(ctx context.Context)       {}
func (c *ClientInformation) AddPlayer(p *server.Player)    {}
func (c *ClientInformation) RemovePlayer(p *server.Player) {}
