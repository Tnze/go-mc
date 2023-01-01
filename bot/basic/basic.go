// Package basic provides some basic packet handler which client needs.
//
// # [Player]
//
// The [Player] is attached to a [Client] by calling [NewPlayer] before the client joins a server.
//
// There is 4 kinds of clientbound packet is handled by this package.
//   - LoginPacket, for cache player info. The player info will be stored in [Player.PlayerInfo].
//   - KeepAlivePacket, for avoid the client to be kicked by the server.
//   - PlayerPosition, is only received when server teleporting the player. And the confirm packet is automatically sent.
//   - Respawn, for updating player info, which may change when player respawned.
//
// # [EventsListener]
//
// Handles some basic event you probably need.
//   - GameStart
//   - ChatMsg
//   - Disconnect
//   - HealthChange
//   - Death
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

func NewPlayer(c *bot.Client, settings Settings, events EventsListener) *Player {
	p := &Player{c: c, Settings: settings}
	c.Events.AddListener(
		bot.PacketHandler{Priority: 0, ID: packetid.ClientboundLogin, F: p.handleLoginPacket},
		bot.PacketHandler{Priority: 0, ID: packetid.ClientboundKeepAlive, F: p.handleKeepAlivePacket},
		bot.PacketHandler{Priority: 0, ID: packetid.ClientboundPlayerPosition, F: p.handlePlayerPosition},
		bot.PacketHandler{Priority: 0, ID: packetid.ClientboundRespawn, F: p.handleRespawnPacket},
	)
	events.attach(p)
	return p
}

func (p *Player) Respawn() error {
	const PerformRespawn = 0

	err := p.c.Conn.WritePacket(pk.Marshal(
		int32(packetid.ServerboundClientCommand),
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
