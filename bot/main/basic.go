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
	*core.Controller
	*screen.Manager
	Settings  basic.Settings
	isSpawn   bool
	fallTicks float32
	jumpTicks float32
}

func NewPlayer(settings basic.Settings) *Player {
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
		Controller: &core.Controller{},
		Manager:    screen.NewManager(),
		Settings:   settings,
		isSpawn:    false,
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
	c.Player.SetPosition(c.Player.Position.Add(c.Player.Motion))

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

	if c.Player.Controller.Jump {
		if c.Player.jumpTicks > 0 {
			c.Player.jumpTicks--
		}
		// TODO: Check if the player is in water or lava
		if c.Player.OnGround && c.Player.jumpTicks == 0 {
			if getBlock, err := c.World.GetBlock(c.Player.Position); !err.Is(basic.NoError) {
				return err
			} else {
				if blockUnder := block.StateList[getBlock]; blockUnder.ID() == (block.HoneyBlock{}).ID() {
					c.Player.Motion = c.Player.Motion.Offset(0, 0.42*core.HoneyBlockMultiplier, 0)
				} else {
					c.Player.Motion = c.Player.Motion.Offset(0, 0.42, 0)
				}

				// TODO: Jump boost
				if c.Player.Controller.Sprint {
					yaw := math.Pi - c.Player.Rotation.Y
					c.Player.Motion = c.Player.Motion.Offset(
						float32(0.2*math.Cos(float64(yaw))),
						0,
						float32(0.2*math.Sin(float64(yaw))),
					)
				}
				c.Player.jumpTicks = 10 // Auto jump cooldown
			}
		} else {
			c.Player.jumpTicks = 0
		}
	}

	strafe := (c.Player.Controller.Right - c.Player.Controller.Left) * 0.98
	forward := (c.Player.Controller.Forward - c.Player.Controller.Back) * 0.98

	if c.Player.Controller.Sneak {
		strafe *= core.SneakSpeed
		forward *= core.SneakSpeed
	}

	if err := moveEntityWithHeading(c, strafe, forward); !err.Is(basic.NoError) {
		return err
	}

	c.Player.SetLastPosition(c.Player.Position)

	return basic.Error{Err: basic.NoError, Info: nil}
}

func moveEntityWithHeading(c *Client, strafe, forward float32) basic.Error {
	gravityMultiplier := float32(1)
	/*if c.Player.Motion.Y <= 0 && c.Player.Effects[core.EffectSlowFalling] > 0 {
		gravityMultiplier = core.SlowFalling
	}*/

	getBlock, err := c.World.GetBlock(c.Player.Position.Offset(0, 0, 0))
	if !err.Is(basic.NoError) {
		return err
	}
	blockUnder := block.StateList[getBlock]

	if !blockUnder.Is(block.Water{}) && !blockUnder.Is(block.Lava{}) {
		acceleration := float32(core.AirBornAcceleration)
		inertia := float32(core.AirBornInertia)

		if !block.IsAir(getBlock) && c.Player.OnGround {
			inertia = float32(core.Slipperiness(getBlock) * core.AirBornInertia)
			acceleration = 0.1 * (0.1627714 / (inertia * inertia * inertia))
		}

		applyHeading(c, strafe, forward, acceleration)
		// Check if on ladder

		if err := moveEntity(c); !err.Is(basic.NoError) {
			return basic.Error{Err: basic.WriterError, Info: fmt.Errorf("failed to move entity: %v", err)}
		}

		// Apply gravity
		/*if c.Player.Effects[core.SlowFalling] > 0 {
			c.Player.Motion.Y += (0.05 * c.Player.Effects[core.SlowFalling] - c.Player.Motion.Y) * 0.2
		} else {}
		*/
		c.Player.Motion.Y -= core.Gravity * gravityMultiplier

		// Apply friction
		c.Player.Motion = c.Player.Motion.OffsetMul(
			inertia,
			core.AirDrag,
			inertia,
		)
	} else {
		// In lava/water
		acceleration := float32(core.LiquidAcceleration)
		inertia := float32(0)
		gravity := float32(core.WaterGravity)
		if blockUnder.Is(block.Water{}) {
			inertia = float32(core.WaterInertia)
		} else {
			inertia = float32(core.LavaInertia)
		}
		horizontalMotion := inertia

		if blockUnder.Is(block.Water{}) {
			// TODO: Depth strider
		}

		applyHeading(c, strafe, forward, acceleration)
		if err := moveEntity(c); !err.Is(basic.NoError) {
			return basic.Error{Err: basic.WriterError, Info: fmt.Errorf("failed to move entity: %v", err)}
		}

		c.Player.Motion = c.Player.Motion.OffsetMul(
			horizontalMotion,
			inertia,
			horizontalMotion,
		)
		c.Player.Motion = c.Player.Motion.Offset(
			0,
			gravity*gravityMultiplier,
			0,
		)

	}

	return basic.Error{Err: basic.NoError, Info: nil}
}

func applyHeading(c *Client, strafe, forward, multipler float32) {
	speed := float32(math.Sqrt(float64(strafe*strafe + forward*forward)))
	speed = multipler / float32(math.Max(float64(speed), 1.0))
	strafe *= speed
	forward *= speed

	yaw := math.Pi - c.Player.Rotation.Y
	c.Player.Motion.X -= strafe*float32(math.Cos(float64(yaw))) + forward*float32(math.Sin(float64(yaw)))
	c.Player.Motion.Z += forward*float32(math.Cos(float64(yaw))) - strafe*float32(math.Sin(float64(yaw)))
}

func moveEntity(c *Client) basic.Error {
	/*oldVelocity := c.Player.Motion

	if c.Player.Controller.Sneak && c.Player.OnGround {
		step := 0.5

		if getBlock, err := c.World.GetBlock(c.Player.Position.Offset(0, -1, 0)); !err.Is(basic.NoError) {

		}
	}*/

	return basic.Error{Err: basic.NoError, Info: nil}
}

func getSurroundingBB(c *Client, bb core.AxisAlignedBB) []core.AxisAlignedBB {
	return nil
}

func (p *Player) Jump(c *Client) error {
	if p.OnGround {
		p.Motion = p.Motion.Add(maths.Vec3d{X: 0, Y: 0.42, Z: 0})
		p.OnGround = false
	}

	return nil
}

func (p *Player) WalkTo(c *Client, pos maths.Vec3d) error {
	path := c.World.PathFind(p.Position, pos)
	for _, v := range path {
		// Set the motion
		fmt.Println(v.Sub(p.Position))
		p.Motion = v.Sub(p.Position)
	}

	return nil
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
