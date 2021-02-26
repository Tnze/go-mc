package basic

import (
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Player struct {
	c        *bot.Client
	Settings Settings

	PlayerInfo
	WorldInfo
}

func NewPlayer(c *bot.Client, settings Settings) *Player {
	b := &Player{c: c, Settings: settings}
	c.Events.AddListener(
		bot.PacketHandler{Priority: 0, ID: packetid.Login, F: b.handleJoinGamePacket},
		bot.PacketHandler{Priority: 0, ID: packetid.KeepAliveClientbound, F: b.handleKeepAlivePacket},
	)
	return b
}

func (p *Player) Respawn() error {
	const PerformRespawn = 0
	return p.c.Conn.WritePacket(pk.Marshal(
		packetid.ClientCommand,
		pk.VarInt(PerformRespawn),
	))
}
