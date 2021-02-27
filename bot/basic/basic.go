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
	isSpawn bool
}

func NewPlayer(c *bot.Client, settings Settings) *Player {
	b := &Player{c: c, Settings: settings}
	c.Events.AddListener(
		bot.PacketHandler{Priority: 0, ID: packetid.Login, F: b.handleJoinGamePacket},
		bot.PacketHandler{Priority: 0, ID: packetid.KeepAliveClientbound, F: b.handleKeepAlivePacket},
		bot.PacketHandler{Priority: 0, ID: packetid.PositionClientbound, F: b.handlePlayerPositionAndLook},
	)
	return b
}

func (p *Player) Respawn() error {
	const PerformRespawn = 0

	err := p.c.Conn.WritePacket(pk.Marshal(
		packetid.ClientCommand,
		pk.VarInt(PerformRespawn),
	))
	if err != nil {
		return Error{err}
	}

	return nil
}

type Error struct {
	Err error
}

func (e Error) Error() string {
	return "bot/basic: " + e.Err.Error()
}

func (e Error) Unwrap() error {
	return e.Err
}
