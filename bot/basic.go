package bot

import (
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/bot/core"
	"github.com/Tnze/go-mc/bot/screen"
	"github.com/Tnze/go-mc/data/packetid"
	. "github.com/Tnze/go-mc/data/slots"
	pk "github.com/Tnze/go-mc/net/packet"
	"math"
)

type Player struct {
	basic.PlayerInfo
	basic.WorldInfo
	*core.EntityPlayer
	*screen.Manager
	c        *Client
	Settings basic.Settings
	isSpawn  bool
}

func NewPlayer(c *Client, settings basic.Settings) *Player {
	return &Player{
		PlayerInfo: basic.PlayerInfo{},
		WorldInfo:  basic.WorldInfo{},
		EntityPlayer: &core.EntityPlayer{
			EntityLiving: &core.EntityLiving{
				Entity: &core.Entity{},
			},
		},
		Manager:  screen.NewManager(),
		c:        c,
		Settings: settings,
		isSpawn:  false,
	}
}

func (p *Player) Respawn() error {
	const PerformRespawn = 0

	err := p.c.Conn.WritePacket(pk.Marshal(
		packetid.SPacketClientCommand,
		pk.VarInt(PerformRespawn),
	))
	if err != nil {
		return Error{err}
	}

	return nil
}

func (p *Player) Chat(msg string) error {
	var (
		message = pk.String(msg[:int(math.Min(float64(len(msg)), 256))])
		//timestamp = pk.Long(time.Now().Unix())
		/*salt             = pk.Long(0)
		signatureLength  = pk.VarInt(0)
		signature        = pk.String("")
		signaturePreview = pk.Boolean(false)*/
	)

	err := p.c.Conn.WritePacket(pk.Marshal(
		packetid.SPacketChatMessage,
		message,
		//timestamp,
		/*salt,
		signatureLength,
		signature,
		signaturePreview,*/
	))
	if err != nil {
		return Error{err}
	}

	return nil
}

func (p *Player) ContainerClick(id int, slot int16, button byte, mode int32, slots ChangedSlots, carried *Slot) error {
	return p.c.Conn.WritePacket(pk.Marshal(
		packetid.CPacketSetContainerSlot,
		pk.UnsignedByte(id),
		pk.VarInt(p.Manager.StateID),
		pk.Short(slot),
		pk.Byte(button),
		pk.VarInt(mode),
		slots,
		carried,
	))
}
