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
					Position: maths.Vec3d[float64]{},
					Motion:   maths.Vec3d[float64]{},
					Rotation: maths.Vec2d[float64]{},
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
	oldPos := c.Player.Position
	c.Player.SetPosition(c.Player.Position.Add(c.Player.Motion))
	if oldPos == c.Player.Position || c.Player.Motion.Length() < 0.0001 {
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

	if c.Player.Controller.Jump {
		if c.Player.jumpTicks > 0 {
			c.Player.jumpTicks--
		}
		// TODO: Check if the player is in water or lava
		if c.Player.OnGround && c.Player.jumpTicks == 0 {
			if getBlock, err := c.World.GetBlock(c.Player.Position.Offset(0, 0.5, 0)); !err.Is(basic.NoError) {
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
						0.2*math.Cos(yaw),
						0,
						0.2*math.Sin(yaw),
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

func moveEntityWithHeading(c *Client, strafe, forward float64) basic.Error {
	gravityMultiplier := 1.0
	/*if c.Player.Motion.Y <= 0 && c.Player.Effects[core.EffectSlowFalling] > 0 {
		gravityMultiplier = core.SlowFalling
	}*/

	getBlock, err := c.World.GetBlock(c.Player.Position.Offset(0.0, 0.5, 0.0))
	if !err.Is(basic.NoError) {
		return err
	}
	blockUnder := block.StateList[getBlock]

	if !blockUnder.Is(block.Water{}) && !blockUnder.Is(block.Lava{}) {
		acceleration := core.AirBornAcceleration
		inertia := core.AirBornInertia

		/*if !block.IsAir(getBlock) && c.Player.OnGround {
			inertia = float32(core.Slipperiness(getBlock) * core.AirBornInertia)
			acceleration = 0.1 * (0.1627714 / (inertia * inertia * inertia))
		}*/

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
		lastY := c.Player.Motion.Y
		acceleration := core.LiquidAcceleration
		inertia := 0.0
		gravity := core.WaterGravity
		if blockUnder.Is(block.Water{}) {
			inertia = core.WaterInertia
			// TODO: Depth strider
		} else {
			inertia = core.LavaInertia
			gravity = core.LavaGravity
		}
		horizontalMotion := inertia

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

		if len(getSurroundingBB(c, c.Player.BoundingBox.Offset(c.Player.Motion.X, c.Player.Motion.Y+0.6-c.Player.Position.Y+lastY, c.Player.Motion.Z))) == 0 {
			c.Player.Motion.Y = core.OutOfLiquidImpulse
		}
	}

	return basic.Error{Err: basic.NoError, Info: nil}
}

func applyHeading(c *Client, strafe, forward, multiplier float64) {
	speed := math.Sqrt(strafe*strafe + forward*forward)
	speed = multiplier / math.Max(speed, 1.0)
	strafe *= speed
	forward *= speed

	yaw := math.Pi - c.Player.Rotation.Y
	c.Player.Motion.X -= strafe*math.Cos(yaw) + forward*math.Sin(yaw)
	c.Player.Motion.Z += forward*math.Cos(yaw) - strafe*math.Sin(yaw)
}

func moveEntity(c *Client) basic.Error {
	if getBlock, err := c.World.GetBlock(c.Player.Position); !err.Is(basic.NoError) {
		return err
	} else if block.StateList[getBlock].Is(block.Cobweb{}) {
		c.Player.Motion = c.Player.Motion.OffsetMul(0.25, 0.05, 0.25)
	}
	dX, dY, dZ := c.Player.Motion.X, c.Player.Motion.Y, c.Player.Motion.Z
	oDx, oDy, oDz := c.Player.Motion.X, c.Player.Motion.Y, c.Player.Motion.Z

	if c.Player.Controller.Sneak && c.Player.OnGround {
		step := 0.5

		for ; dX != 0 && len(getSurroundingBB(c, c.Player.BoundingBox.Offset(dX, 0, 0))) == 0; oDx = dX {
			if dX < step && dX >= -step {
				dX = 0
			} else if dX > 0 {
				dX -= step
			} else {
				dX += step
			}
		}

		for ; dY != 0 && len(getSurroundingBB(c, c.Player.BoundingBox.Offset(0, 0, dZ))) == 0; oDz = dZ {
			if dZ < step && dZ >= -step {
				dZ = 0
			} else if dZ > 0 {
				dZ -= step
			} else {
				dZ += step
			}
		}

		for dX != 0 && dZ != 0 && len(getSurroundingBB(c, c.Player.BoundingBox.Offset(dX, 0, dZ))) == 0 {
			if dX < step && dX >= -step {
				dX = 0
			} else if dX > 0 {
				dX -= step
			} else {
				dX += step
			}

			if dZ < step && dZ >= -step {
				dZ = 0
			} else if dZ > 0 {
				dZ -= step
			} else {
				dZ += step
			}

			oDx = dX
			oDz = dZ
		}
	}

	playerBB := c.Player.BoundingBox
	queryBB := playerBB.Offset(dX, dY, dZ)
	collidingBB := getSurroundingBB(c, queryBB)
	oldBB := c.Player.BoundingBox

	for _, bb := range collidingBB {
		dY = bb.CollideY(playerBB, dY)
	}
	playerBB = playerBB.Offset(0, dY, 0)

	for _, bb := range collidingBB {
		dX = bb.CollideX(playerBB, dX)
	}
	playerBB = playerBB.Offset(dX, 0, 0)

	for _, bb := range collidingBB {
		dZ = bb.CollideZ(playerBB, dZ)
	}
	playerBB = playerBB.Offset(0, 0, dZ)

	if c.Player.OnGround || (oDy != dY && oDy < 0) && (dX != oDx || dZ != oDz) {
		oVXC, oVYC, oVZC := dX, dY, dZ

		// Step up blocks
		dY = core.StepHeight
		queryBB = oldBB.Expand(c.Player.Motion.X, dY, oDz)
		collidingBB = getSurroundingBB(c, queryBB)

		BB1, BB2 := oldBB, oldBB
		BB_XZ := BB1.Expand(dX, 0, dZ)

		dY1, dY2 := oDy, oDy
		for _, bb := range collidingBB {
			dY1 = bb.CollideY(BB_XZ, dY1)
			dY2 = bb.CollideY(BB1, dY2)
		}
		BB1 = BB1.Offset(0, dY1, 0)
		BB2 = BB2.Offset(0, dY2, 0)

		dX1, dX2 := oDx, oDx
		for _, bb := range collidingBB {
			dX1 = bb.CollideX(BB1, dX1)
			dX2 = bb.CollideX(BB2, dX2)
		}
		BB1 = BB1.Offset(dX1, 0, 0)
		BB2 = BB2.Offset(dX2, 0, 0)

		dZ1, dZ2 := dZ, dZ
		for _, bb := range collidingBB {
			dZ1 = bb.CollideZ(BB1, dZ1)
			dZ2 = bb.CollideZ(BB2, dZ2)
		}
		BB1 = BB1.Offset(0, 0, dZ1)
		BB2 = BB2.Offset(0, 0, dZ2)

		norm1, norm2 := dX1*dX1+dZ1*dZ1, dX2*dX2+dZ2*dZ2

		if norm1 > norm2 {
			dX, dY, dZ = dX1, -dY1, dZ1
		} else {
			dX, dY, dZ = dX2, -dY2, dZ2
			playerBB = BB2
		}

		for _, bb := range collidingBB {
			dY = bb.CollideY(playerBB, dY)
		}
		playerBB = playerBB.Offset(0, dY, 0)

		if oVXC*oVXC+oVZC*oVZC >= dX*dX+dZ*dZ {
			dX, dY, dZ = oVXC, oVYC, oVZC
			playerBB = oldBB
		}
	} else {
		c.Player.OnGround = false
	}

	c.Player.Position = c.Player.Position.Offset(
		playerBB.MinX+0.3,
		playerBB.MinY,
		playerBB.MinZ+0.3,
	)

	// Apply block collision
	playerBB = playerBB.Contract(0.001, 0.001, 0.001)
	cursor := maths.NullVec3d
	for cursor.Y = math.Floor(playerBB.MinY); cursor.Y < math.Ceil(playerBB.MaxY); cursor.Y++ {
		for cursor.Z = math.Floor(playerBB.MinZ); cursor.Z < math.Ceil(playerBB.MaxZ); cursor.Z++ {
			for cursor.X = math.Floor(playerBB.MinX); cursor.X < math.Ceil(playerBB.MaxX); cursor.X++ {
				if getBlock, err := c.World.GetBlock(cursor); !err.Is(basic.NoError) {
					continue
				} else {
					if block.StateList[getBlock].Is(block.SoulSand{}) {
						c.Player.Motion.X *= core.SoulSandMultiplier
						c.Player.Motion.Z *= core.SoulSandMultiplier
					} else if block.StateList[getBlock].Is(block.HoneyBlock{}) {
						c.Player.Motion.X *= core.HoneyBlockMultiplier
						c.Player.Motion.Z *= core.HoneyBlockMultiplier
					} else if block.StateList[getBlock].Is(block.Cobweb{}) {
						// Set cobweb to true
					}
				}
			}
		}
	}

	if blockBelow, err := c.World.GetBlock(c.Player.Position.Offset(0, -0.5, 0)); !err.Is(basic.NoError) {
		return err
	} else {
		if block.StateList[blockBelow].Is(block.SoulSand{}) {
			c.Player.Motion.X *= core.SoulSandMultiplier
			c.Player.Motion.Z *= core.SoulSandMultiplier
		} else if block.StateList[blockBelow].Is(block.HoneyBlock{}) {
			c.Player.Motion.X *= core.HoneyBlockMultiplier
			c.Player.Motion.Z *= core.HoneyBlockMultiplier
		} else if block.StateList[blockBelow].Is(block.Cobweb{}) {
			// Set cobweb to true
		}
	}

	return basic.Error{Err: basic.NoError, Info: nil}
}

func getSurroundingBB(c *Client, bb core.AxisAlignedBB[float64]) []core.AxisAlignedBB[float64] {
	var blocks []core.AxisAlignedBB[float64]
	for y := bb.MinY; y < bb.MaxY; y++ {
		for z := bb.MinZ; z < bb.MaxZ; z++ {
			for x := bb.MinX; x < bb.MaxX; x++ {
				if getBlock, err := c.World.GetBlock(maths.Vec3d[float64]{X: x, Y: y, Z: z}); !err.Is(basic.NoError) {
					continue
				} else {
					if block.StateList[getBlock].IsAir() {
						continue
					}

					blocks = append(blocks, block.StateList[getBlock].GetCollisionBox())
				}
			}
		}
	}

	return blocks
}

func (p *Player) Jump(c *Client) error {
	if p.OnGround {
		p.Motion = p.Motion.Add(maths.Vec3d[float64]{X: 0, Y: 0.42, Z: 0})
		p.OnGround = false
	}

	return nil
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
