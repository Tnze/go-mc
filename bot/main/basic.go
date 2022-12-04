package main

import (
	"fmt"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/bot/core"
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/bot/screen"
	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/data/packetid"
	. "github.com/Tnze/go-mc/data/slots"
	"github.com/Tnze/go-mc/level/block"
	pk "github.com/Tnze/go-mc/net/packet"
	"math"
)

type Player struct {
	world.PlayerInfo
	world.WorldInfo
	*core.EntityPlayer
	*screen.Manager
	Settings  basic.Settings
	isSpawn   bool
	fallTicks float32
}

func NewPlayer(c *Client, settings basic.Settings) *Player {
	return &Player{
		PlayerInfo: world.PlayerInfo{},
		WorldInfo:  world.WorldInfo{},
		EntityPlayer: &core.EntityPlayer{
			EntityLiving: &core.EntityLiving{
				Entity: &core.Entity{
					Position: maths.Vec3d{},
					Motion:   maths.Vec3d{},
					Rotation: maths.Vec2d{},
				},
			},
		},
		Manager:  screen.NewManager(),
		Settings: settings,
		isSpawn:  false,
	}
}

func (p *Player) Respawn(c *Client) error {
	const PerformRespawn = 0

	err := c.Conn.WritePacket(pk.Marshal(
		packetid.SPacketClientCommand,
		pk.VarInt(PerformRespawn),
	))
	if !err.Is(basic.NoError) {
		return basic.Error{Err: basic.WriterError, Info: err}
	}

	return nil
}

func (p *Player) Chat(c *Client, msg string) error {
	var (
		message = pk.String(msg[:int(math.Min(float64(len(msg)), 256))])
		//timestamp = pk.Long(time.Now().Unix())
		/*salt             = pk.Long(0)
		signatureLength  = pk.VarInt(0)
		signature        = pk.String("")
		signaturePreview = pk.Boolean(false)*/
	)

	err := c.Conn.WritePacket(pk.Marshal(
		packetid.SPacketChatMessage,
		message,
		//timestamp,
		/*salt,
		signatureLength,
		signature,
		signaturePreview,*/
	))
	if !err.Is(basic.NoError) {
		return basic.Error{Err: basic.WriterError, Info: err}
	}

	return nil
}

func ApplyPhysics(c *Client) basic.Error {
	c.Player.Position = c.Player.Position.Add(c.Player.Motion)
	if c.Player.Position.DistanceTo(c.Player.GetLastPosition()) < 0.01 {
		c.Player.Motion = maths.NullVec3d
		return basic.Error{Err: basic.NoError, Info: nil}
	}
	if err := c.Conn.WritePacket(
		pk.Marshal(
			packetid.SPacketPlayerPosition,
			pk.Double(c.Player.Position.X),
			pk.Double(c.Player.Position.Y),
			pk.Double(c.Player.Position.Z),
			pk.Boolean(c.Player.OnGround),
		),
	); !err.Is(basic.NoError) {
		return basic.Error{Err: basic.WriterError, Info: fmt.Errorf("failed to send player position: %v", err)}
	}

	// Apply inertia
	getBlock, err := c.World.GetBlock(c.Player.Position)
	if !err.Is(basic.NoError) {
		return err
	}

	inertia := float32(core.Slipperiness(getBlock) * core.AirBornInertia)
	c.Player.Motion = c.Player.Motion.MulScalar(inertia)

	// Check Y motion
	if c.Player.Motion.Y != 0 {
		if getBlock, err := c.World.GetBlock(c.Player.Position.Sub(maths.Vec3d{Y: 0.5})); !err.Is(basic.NoError) {
			return err
		} else if block.IsAir(getBlock) {
			c.Player.fallTicks++
		} else {
			c.Player.fallTicks = 0
		}
	}

	// Apply gravity
	c.Player.Motion = c.Player.Motion.Add(maths.Vec3d{X: 0, Y: -core.Gravity * (c.Player.fallTicks + 1), Z: 0})

	// Apply friction
	/*c.Player.Motion.X *= inertia
	c.Player.Motion.Y *= core.AirDrag
	c.Player.Motion.Z *= inertia*/

	// Reset motion if it's too small
	if c.Player.Motion.Length() < 0.05 {
		c.Player.Motion = maths.NullVec3d
	}

	c.Player.SetLastPosition(c.Player.Position)

	return basic.Error{Err: basic.NoError, Info: nil}
}

func (p *Player) ContainerClick(c *Client, id int, slot int16, button byte, mode int32, slots ChangedSlots, carried *Slot) error {
	return c.Conn.WritePacket(pk.Marshal(
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
