// Package phy implements a minimal physics simulation necessary for realistic
// bot behavior.
package phy

import (
	"math"

	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/bot/world/entity/player"
)

const (
	playerWidth  = 0.6
	playerHeight = 1.8
	resetVel     = 0.003

	maxYawChange   = 33
	maxPitchChange = 22

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

// Point represents a point in 3D space.
type Point struct {
	X, Y, Z float64
}

// State tracks physics state.
type State struct {
	// player state.
	Pos        Point
	Vel        Point
	Yaw, Pitch float64

	// player state flags.
	onGround  bool
	collision struct {
		vertical   bool
		horizontal bool
	}

	Run bool
}

func (s *State) ServerPositionUpdate(player player.Pos, w World) error {
	s.Pos = Point{X: player.X, Y: player.Y, Z: player.Z}
	s.Yaw, s.Pitch = float64(player.Yaw), float64(player.Pitch)
	s.Vel = Point{}
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
				if block := w.GetBlockStatus(x, y, z); block > 0 {
					out = append(out, AABB{X: MinMax{Max: 1}, Y: MinMax{Max: 1}, Z: MinMax{Max: 1}, Block: block}.Offset(float64(x), float64(y), float64(z)))
				}
			}
		}
	}
	return out
}

func (s *State) applyLookInputs(input Inputs) {
	errYaw := math.Min(math.Max(input.Yaw-s.Yaw, -maxYawChange), maxYawChange)
	s.Yaw += errYaw
}

func (s *State) applyPosInputs(input Inputs, acceleration, inertia float64) {
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

func (s *State) Tick(input Inputs, w World) error {
	if !s.Run {
		return nil
	}
	var inertia = inertia
	var acceleration = acceleration
	if s.onGround {
		inertia *= slipperiness
		acceleration = 0.1 * (0.1627714 / (inertia * inertia * inertia))
	}
	s.applyLookInputs(input)
	s.applyPosInputs(input, acceleration, inertia)

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
	// Gravity
	s.Vel.Y -= gravity
	// Drag & friction.
	s.Vel.Y *= drag
	s.Vel.X *= inertia
	s.Vel.Z *= inertia

	// Apply collision.
	var (
		player       = s.BB()
		query        = player.Extend(s.Vel.X, s.Vel.Y, s.Vel.Z)
		surroundings = s.surroundings(query, w)
		newVel       = s.Vel
	)
	for _, b := range surroundings {
		newVel.Y = b.YOffset(player, newVel.Y)
	}
	player = player.Offset(0, newVel.Y, 0)
	for _, b := range surroundings {
		newVel.X = b.XOffset(player, newVel.X)
	}
	player = player.Offset(newVel.X, 0, 0)
	for _, b := range surroundings {
		newVel.Z = b.ZOffset(player, newVel.Z)
	}
	player = player.Offset(0, 0, newVel.Z)

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
