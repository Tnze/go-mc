// Package phy implements a minimal physics simulation necessary for realistic
// bot behavior.
package phy

import (
	"math"

	"github.com/Tnze/go-mc/bot/path"
	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/bot/world/entity/player"
)

const (
	playerWidth  = 0.6
	playerHeight = 1.8
	resetVel     = 0.003

	maxYawChange   = 33
	maxPitchChange = 11

	stepHeight   = 0.6
	minJumpTicks = 14

	gravity      = 0.08
	drag         = 0.98
	acceleration = 0.02
	inertia      = 0.91
	slipperiness = 0.6
)

// World represents a provider of information about the surrounding world.
type World interface {
	GetBlockStatus(x, y, z int) world.BlockStatus
}

// Surrounds represents the blocks surrounding the player (Y, Z, X).
type Surrounds []AABB

// State tracks physics state.
type State struct {
	// player state.
	Pos        path.Point
	Vel        path.Point
	Yaw, Pitch float64
	lastJump   uint32

	// player state flags.
	onGround  bool
	collision struct {
		vertical   bool
		horizontal bool
	}

	tick uint32
	Run  bool
}

func (s *State) ServerPositionUpdate(player player.Pos, w World) error {
	s.Pos = path.Point{X: player.X, Y: player.Y, Z: player.Z}
	s.Yaw, s.Pitch = float64(player.Yaw), float64(player.Pitch)
	s.Vel = path.Point{}
	s.onGround, s.collision.vertical, s.collision.horizontal = false, false, false
	s.Run = true
	return nil
}

func abs(i1, i2 int) int {
	if i1 < i2 {
		return i2 - i1
	}
	return i1 - i2
}

func (s *State) surroundings(query AABB, w World) Surrounds {
	minY, maxY := int(math.Floor(query.Y.Min))-1, int(math.Floor(query.Y.Max))+1
	minZ, maxZ := int(math.Floor(query.Z.Min)), int(math.Floor(query.Z.Max))+1
	minX, maxX := int(math.Floor(query.X.Min)), int(math.Floor(query.X.Max))+1

	out := Surrounds(make([]AABB, 0, abs(maxY, minY)*abs(maxZ, minZ)*abs(maxX, minX)))
	for y := minY; y < maxY; y++ {
		for z := minZ; z < maxZ; z++ {
			for x := minX; x < maxX; x++ {
				if block := w.GetBlockStatus(x, y, z); !path.AirLikeBlock(block) {
					out = append(out, AABB{X: MinMax{Max: 1}, Y: MinMax{Max: 1}, Z: MinMax{Max: 1}, Block: block}.Offset(float64(x), float64(y), float64(z)))
				}
			}
		}
	}
	return out
}

func (s *State) BB() AABB {
	return AABB{
		X: MinMax{Min: -playerWidth / 2, Max: playerWidth / 2},
		Y: MinMax{Max: playerHeight},
		Z: MinMax{Min: -playerWidth / 2, Max: playerWidth / 2},
	}.Offset(s.Pos.X, s.Pos.Y, s.Pos.Z)
}

func (s *State) Position() player.Pos {
	return player.Pos{
		X: s.Pos.X, Y: s.Pos.Y, Z: s.Pos.Z,
		Yaw: float32(s.Yaw), Pitch: float32(s.Pitch),
		OnGround: s.onGround,
	}
}

func (s *State) Tick(input path.Inputs, w World) error {
	s.tick++
	if !s.Run {
		return nil
	}
	s.tickVelocity(input, w)

	player, newVel := s.computeCollision(s.BB(), s.BB().Extend(s.Vel.X, s.Vel.Y, s.Vel.Z), w)

	bb := player.Extend(s.Vel.X, stepHeight, s.Vel.Z)
	surroundings := s.surroundings(bb, w)
	y := float64(0)
	for _, b := range surroundings {
		if b.Intersects(bb) && bb.Y.Max > b.Y.Min {
			y = math.Max(y, b.Y.Max)
		}
	}
	//fmt.Printf("pY = %.2f, maxblockY = %.1f (delta = %.1f)\n", bb.Y.Min, y, bb.Y.Min-y)
	if d := bb.Y.Min - y; d >= -stepHeight && d < stepHeight-1 {
		bb := player.Offset(0, stepHeight, 0)
		player, newVel = s.computeCollision(bb, bb.Extend(s.Vel.X, s.Vel.Y, s.Vel.Z), w)
	}

	// Update flags.
	s.Pos.X = player.X.Min + playerWidth/2
	s.Pos.Y = player.Y.Min
	s.Pos.Z = player.Z.Min + playerWidth/2
	s.collision.horizontal = newVel.X != s.Vel.X || newVel.Z != s.Vel.Z
	s.collision.vertical = newVel.Y != s.Vel.Y
	s.onGround = s.collision.vertical && s.Vel.Y < 0

	s.Vel = newVel
	return nil
}

func (s *State) applyLookInputs(input path.Inputs) {
	errYaw := math.Min(math.Max(input.Yaw-s.Yaw, -maxYawChange), maxYawChange)
	s.Yaw += errYaw
	errPitch := math.Min(math.Max(input.Pitch-s.Pitch, -maxPitchChange), maxPitchChange)
	s.Pitch += errPitch
}

func (s *State) applyPosInputs(input path.Inputs, acceleration, inertia float64) {
	// fmt.Println(input.Jump, s.lastJump, s.onGround)
	if input.Jump && s.lastJump+minJumpTicks < s.tick {
		s.lastJump = s.tick
		s.Vel.Y += 0.42
	}

	speed := math.Sqrt(input.ThrottleX*input.ThrottleX + input.ThrottleZ*input.ThrottleZ)
	if speed < 0.01 {
		return
	}
	speed = acceleration / math.Max(speed, 1)

	input.ThrottleX *= speed
	input.ThrottleZ *= speed

	s.Vel.X += input.ThrottleX
	s.Vel.Z += input.ThrottleZ
}

func (s *State) tickVelocity(input path.Inputs, w World) {
	var inertia = inertia
	var acceleration = acceleration
	if below := w.GetBlockStatus(int(math.Floor(s.Pos.X)), int(math.Floor(s.Pos.Y))-1, int(math.Floor(s.Pos.Z))); s.onGround && !path.AirLikeBlock(below) {
		inertia *= slipperiness
		acceleration = 0.1 * (0.1627714 / (inertia * inertia * inertia))
	}

	// Deadzone velocities when they get too low.
	if math.Abs(s.Vel.X) < resetVel {
		s.Vel.X = 0
	}
	if math.Abs(s.Vel.Y) < resetVel {
		s.Vel.Y = 0
	}
	if math.Abs(s.Vel.Z) < resetVel {
		s.Vel.Z = 0
	}

	s.applyLookInputs(input)
	s.applyPosInputs(input, acceleration, inertia)

	// Gravity
	s.Vel.Y -= gravity
	// Drag & friction.
	s.Vel.Y *= drag
	s.Vel.X *= inertia
	s.Vel.Z *= inertia
}

func (s *State) computeCollision(bb, query AABB, w World) (outBB AABB, outVel path.Point) {
	surroundings := s.surroundings(query, w)
	outVel = s.Vel

	for _, b := range surroundings {
		outVel.Y = b.YOffset(bb, outVel.Y)
	}
	bb = bb.Offset(0, outVel.Y, 0)
	for _, b := range surroundings {
		outVel.X = b.XOffset(bb, outVel.X)
	}
	bb = bb.Offset(outVel.X, 0, 0)
	for _, b := range surroundings {
		outVel.Z = b.ZOffset(bb, outVel.Z)
	}
	bb = bb.Offset(0, 0, outVel.Z)
	return bb, outVel
}
