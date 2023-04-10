package main

import (
	"fmt"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/bot/core"
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/bot/screen"
	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/data/effects"
	"github.com/Tnze/go-mc/data/enums"
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
	Settings             basic.Settings
	isSpawn              bool
	fallTicks            float32
	fallDistance         float32
	stepHeight           float32
	jumpTicks            float32
	collidedHorizontally bool
	collidedVertically   bool
	collided             bool
}

func NewPlayer(settings basic.Settings) *Player {
	return &Player{
		PlayerInfo: world.PlayerInfo{},
		WorldInfo:  world.WorldInfo{},
		EntityPlayer: &core.EntityPlayer{
			EntityLiving: &core.EntityLiving{
				Entity: &core.Entity{
					Position: maths.Vec3d[float64]{},
					Motion:   maths.Vec3d[float64]{},
					Rotation: maths.Vec2d[float64]{},
				},
				ActivePotionEffects: make(map[int32]*effects.EffectStatus),
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

func (p *Player) handleJumpWater() {
	p.Motion.Y += 0.03999999910593033
}

func (p *Player) handleJumpLava() {
	p.Motion.Y += 0.03999999910593033
}

func (p *Player) Jump() {
	p.Motion.Y += 0.42

	if p.IsPotionActive(effects.JumpBoost) {
		p.Motion.Y += 0.1 * float64(p.GetPotionEffect(effects.JumpBoost).Amplifier+1)
	}

	// TODO: Check if the player is sprinting
}

func (p *Player) travel(c *Client) error {
	if getBlock, err := c.World.GetBlock(p.Position); !err.Is(basic.NoError) {
		return err
	} else {
		if !getBlock.Is(block.Water{}) {
			if !getBlock.Is(block.Lava{}) {
				// Replace with elytra flying
				if false {
					if p.Motion.Y > -0.5 {
						p.fallDistance = 1.0
					}

					vecRotation := maths.GetVectorFromRotation(p.Rotation)
					f := p.Rotation.X * 0.017453292
					d6 := math.Sqrt(vecRotation.X*vecRotation.X + vecRotation.Z*vecRotation.Z)
					d8 := math.Sqrt(p.Motion.X*p.Motion.X + p.Motion.Z*p.Motion.Z)
					d1 := vecRotation.Length()
					f4 := math.Cos(f)
					f4 = f4 * f4 * math.Min(1.0, d1/0.4)
					p.Motion.Y += -0.08 + f4*0.06

					if p.Motion.Y < 0.0 && d6 > 0.0 {
						d2 := p.Motion.Y * -0.1 * f4
						p.Motion.Y += d2
						p.Motion.X += vecRotation.X * d2 / d6
						p.Motion.Z += vecRotation.Z * d2 / d6
					}

					if f < 0.0 {
						d10 := d8 * -math.Sin(f) * 0.04
						p.Motion.Y += d10 * 3.2
						p.Motion.X -= vecRotation.X * d10 / d6
						p.Motion.Z -= vecRotation.Z * d10 / d6
					}

					if d6 > 0.0 {
						p.Motion.X += (vecRotation.X/d6*d8 - p.Motion.X) * 0.1
						p.Motion.Z += (vecRotation.Z/d6*d8 - p.Motion.Z) * 0.1
					}

					p.Motion.X *= 0.9900000095367432
					p.Motion.Y *= 0.9800000190734863
					p.Motion.Z *= 0.9900000095367432
					p.move(enums.MoverTypeSelf, c, p.Motion)
				} else {
					f6 := 0.91

					if p.OnGround {
						position := maths.Vec3d[float64]{X: p.Position.X, Y: p.BoundingBox.MinY - 1, Z: p.Position.Z}
						getBlock, _ := c.World.GetBlock(position)
						f6 = enums.Slipperiness(getBlock.StateID())
					}

					//f7 := 0.16277136 / (f6 * f6 * f6)
					f8 := 0.02
					p.moveRelative(p.Motion, f8)
					f6 = 0.91

					if p.OnGround {
						position := maths.Vec3d[float64]{X: p.Position.X, Y: p.BoundingBox.MinY - 1, Z: p.Position.Z}
						getBlock, _ := c.World.GetBlock(position)
						f6 = enums.Slipperiness(getBlock.StateID())
					}

					// TODO: Check if the player is on a ladder

					p.move(enums.MoverTypeSelf, c, p.Motion)

					// TODO: Check if the player is on a ladder and collided horizontally

					if p.IsPotionActive(effects.Levitation) {
						p.Motion.Y += (0.05*float64(p.GetPotionEffect(effects.Levitation).Amplifier+1) - p.Motion.Y) * 0.2
					} else {
						p.Motion.Y -= 0.08
					}

					p.Motion.Y *= 0.9800000190734863
					p.Motion.X *= f6
					p.Motion.Z *= f6
				}
			} else {
				//d4 := p.Position.Y
				p.moveRelative(p.Motion, 0.02)
				p.move(enums.MoverTypeSelf, c, p.Motion)

				p.Motion.X *= 0.5
				p.Motion.Y *= 0.5
				p.Motion.Z *= 0.5

				p.Motion.Y -= 0.02

				if p.collidedHorizontally /*&& p.isOffsetPositionInLiquid(p.Motion.X, p.Motion.Y+0.6000000238418579-p.Position.Y+d4, p.Motion.Z)*/ {
					p.Motion.Y = 0.30000001192092896
				}
			}
		} else {
			//d0 := p.Position.Y
			f1 := enums.WaterInertia
			f2 := 0.02

			// TODO: Get depth strider enchantment level

			p.moveRelative(p.Motion, f2)
			p.move(enums.MoverTypeSelf, c, p.Motion)
			p.Motion.X *= f1
			p.Motion.Y *= 0.800000011920929
			p.Motion.Z *= f1

			p.Motion.Y -= 0.02

			if p.collidedHorizontally /*&& p.isOffsetPositionInLiquid(p.Motion.X, p.Motion.Y+0.6000000238418579-p.Position.Y+d0, p.Motion.Z)*/ {
				p.Motion.Y = 0.30000001192092896
			}
		}
	}

	return nil
}

func (p *Player) move(moveType enums.MoverType, c *Client, motion maths.Vec3d[float64]) error {
	if moveType == enums.MoverTypePiston {
	}

	x := motion.X
	y := motion.Y
	z := motion.Z

	if getBlock, err := c.World.GetBlock(p.Position); !err.Is(basic.NoError) {
		return err
	} else {
		if getBlock.Is(block.Cobweb{}) {
			x *= 0.25
			y *= 0.05000000074505806
			z *= 0.25
			p.Motion = maths.NullVec3d
		}
	}

	if (moveType == enums.MoverTypePlayer || moveType == enums.MoverTypeSelf) && p.OnGround {
		/*for ; x != 0.0 && len(c.World.GetCollisionBoxes(*p.Entity, p.BoundingBox.Offset(x, -1.0, 0.0))) == 0; x = motion.X {
			if x < 0.05 && x >= -0.05 {
				x = 0.0
			} else if x > 0.0 {
				x -= 0.05
			} else {
				x += 0.05
			}
		}

		for ; z != 0.0 && len(c.World.GetCollisionBoxes(*p.Entity, p.BoundingBox.Offset(0.0, -1.0, z))) == 0; y = motion.Y {
			if z < 0.05 && z >= -0.05 {
				z = 0.0
			} else if z > 0.0 {
				z -= 0.05
			} else {
				z += 0.05
			}
		}

		for ; x != 0.0 && z != 0.0 && len(c.World.GetCollisionBoxes(*p.Entity, p.BoundingBox.Offset(x, -1.0, z))) == 0; z = motion.Z {
			if x < 0.05 && x >= -0.05 {
				x = 0.0
			} else if x > 0.0 {
				x -= 0.05
			} else {
				x += 0.05
			}

			x = motion.X

			if z < 0.05 && z >= -0.05 {
				z = 0.0
			} else if z > 0.0 {
				z -= 0.05
			} else {
				z += 0.05
			}
		}*/

		// TODO: Add entity collision

		//flag := p.OnGround || y != motion.Y && y < 0.0

		// TODO: Add step

		p.collidedHorizontally = x != motion.X || z != motion.Z
		p.collidedVertically = y != motion.Y
		p.OnGround = p.collidedVertically && y < 0.0
		p.collided = p.collidedHorizontally || p.collidedVertically
		//j6 := math.Floor(p.Position.X)
		//i1 := math.Floor(p.Position.Y - 0.20000000298023224)
		//k6 := math.Floor(p.Position.Z)
		//getBlock, _ := c.World.GetBlock(maths.Vec3d[float64]{X: j6, Y: i1, Z: k6})
		if p.OnGround {
			p.fallDistance = 0.0
		} else if y < 0.0 {
			p.fallDistance -= float32(y)
		}

		if x != motion.X {
			p.Motion.X = 0.0
		}

		if y != motion.Y {
			p.Motion.Y = 0.0
		}

		if z != motion.Z {
			p.Motion.Z = 0.0
		}
	}
	return nil
}

func (p *Player) moveRelative(motion maths.Vec3d[float64], friction float64) {
}

func (p *Player) updateFallState(y float64, onGround bool, getBlock block.Block) {
	if onGround {
		p.fallDistance = 0.0
	} else if y < 0.0 {
		p.fallDistance -= float32(y)
	}
}

func ApplyPhysics(c *Client) basic.Error {
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

	if c.Player.jumpTicks > 0 {
		c.Player.jumpTicks--
	}

	/*if c.Player.Motion.Length() < enums.NegligibleVelocity {
		c.Player.Motion = maths.NullVec3d
		return basic.Error{Err: basic.NoError, Info: nil}
	}*/

	if c.Player.Controller.Jump {
		if getBlock, err := c.World.GetBlock(c.Player.Position); !err.Is(basic.NoError) {
			return err
		} else {
			if getBlock.Is(block.Water{}) {
				c.Player.handleJumpWater()
			} else if getBlock.Is(block.Lava{}) {
				c.Player.handleJumpLava()
			} else if c.Player.OnGround && c.Player.jumpTicks == 0 {
				c.Player.Jump()
				c.Player.jumpTicks = 10
			}
		}
	} else {
		c.Player.jumpTicks = 0
	}

	// Start section: Travel
	c.Player.travel(c)

	c.Player.Position = c.Player.Position.Add(c.Player.Motion)

	return basic.Error{Err: basic.NoError, Info: nil}
}

func (p *Player) WalkTo(c *Client, pos maths.Vec3d[float64]) error {
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
