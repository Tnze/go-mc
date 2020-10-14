// Package phy implements a minimal physics simulation necessary for realistic
// bot behavior.
package phy

import (
	"fmt"
	"math"

	"github.com/Tnze/go-mc/bot/path"
	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/bot/world/entity/player"
	"github.com/Tnze/go-mc/data/block/shape"
)

const (
	playerWidth  = 0.6
	playerHeight = 1.8
	resetVel     = 0.003

	maxYawChange   = 11
	maxPitchChange = 7

	stepHeight       = 0.6
	minJumpTicks     = 14
	ladderMaxSpeed   = 0.15
	ladderClimbSpeed = 0.2

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
	fmt.Printf("TELEPORT (y=%0.2f, velY=%0.3f): %0.2f, %0.2f, %0.2f\n", s.Pos.Y, s.Vel.Y, player.X-s.Pos.X, player.Y-s.Pos.Y, player.Z-s.Pos.Z)

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

	out := Surrounds(make([]AABB, 0, abs(maxY, minY)*abs(maxZ, minZ)*abs(maxX, minX)*2))
	for y := minY; y < maxY; y++ {
		for z := minZ; z < maxZ; z++ {
			for x := minX; x < maxX; x++ {
				bStateID := w.GetBlockStatus(x, y, z)
				if !path.AirLikeBlock(bStateID) {
					bbs, err := shape.CollisionBoxes(bStateID)
					if err != nil {
						panic(err)
					}
					for _, box := range bbs {
						out = append(out, AABB{
							X:     MinMax{Min: box.Min.X, Max: box.Max.X},
							Y:     MinMax{Min: box.Min.Y, Max: box.Max.Y},
							Z:     MinMax{Min: box.Min.Z, Max: box.Max.Z},
							Block: bStateID,
						}.Offset(float64(x), float64(y), float64(z)))
					}
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

	var inertia = inertia
	var acceleration = acceleration
	if below := w.GetBlockStatus(int(math.Floor(s.Pos.X)), int(math.Floor(s.Pos.Y))-1, int(math.Floor(s.Pos.Z))); s.onGround && !path.AirLikeBlock(below) {
		inertia *= slipperiness
		acceleration = 0.1 * (0.1627714 / (inertia * inertia * inertia))
	}

	s.tickVelocity(input, inertia, acceleration, w)
	s.tickPosition(w)

	if path.IsLadder(w.GetBlockStatus(int(math.Floor(s.Pos.X)), int(math.Floor(s.Pos.Y)), int(math.Floor(s.Pos.Z)))) && s.collision.horizontal {
		s.Vel.Y = ladderClimbSpeed
	}

	// Gravity
	s.Vel.Y -= gravity
	// Drag & friction.
	s.Vel.Y *= drag
	s.Vel.X *= inertia
	s.Vel.Z *= inertia
	return nil
}

func (s *State) tickVelocity(input path.Inputs, inertia, acceleration float64, w World) {

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

	lower := w.GetBlockStatus(int(math.Floor(s.Pos.X)), int(math.Floor(s.Pos.Y)), int(math.Floor(s.Pos.Z)))
	if path.IsLadder(lower) {
		s.Vel.X = math.Min(math.Max(-ladderMaxSpeed, s.Vel.X), ladderMaxSpeed)
		s.Vel.Z = math.Min(math.Max(-ladderMaxSpeed, s.Vel.Z), ladderMaxSpeed)
		s.Vel.Y = math.Min(math.Max(-ladderMaxSpeed, s.Vel.Y), ladderMaxSpeed)
	}
}

func (s *State) applyLookInputs(input path.Inputs) {
	if !math.IsNaN(input.Yaw) {
		errYaw := math.Min(math.Max(modYaw(input.Yaw, s.Yaw), -maxYawChange), maxYawChange)
		s.Yaw += errYaw
	}
	errPitch := math.Min(math.Max(input.Pitch-s.Pitch, -maxPitchChange), maxPitchChange)
	s.Pitch += errPitch
}

func (s *State) applyPosInputs(input path.Inputs, acceleration, inertia float64) {
	// fmt.Println(input.Jump, s.lastJump, s.onGround)
	if input.Jump && s.lastJump+minJumpTicks < s.tick && s.onGround {
		s.lastJump = s.tick
		s.Vel.Y = 0.42
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

func (s *State) tickPosition(w World) {
	// fmt.Printf("TICK POSITION: %0.2f, %0.2f, %0.2f - (%0.2f, %0.2f, %0.2f)\n", s.Pos.X, s.Pos.Y, s.Pos.Z, s.Vel.X, s.Vel.Y, s.Vel.Z)

	player, newVel := s.computeCollisionYXZ(s.BB(), s.BB().Offset(s.Vel.X, s.Vel.Y, s.Vel.Z), s.Vel, w)
	//fmt.Printf("offset = %0.2f, %0.2f, %0.2f\n", player.X.Min-s.Pos.X, player.Y.Min-s.Pos.Y, player.Z.Min-s.Pos.Z)

	//fmt.Printf("onGround = %v, s.Vel.Y = %0.3f, newVel.Y = %0.3f\n", s.onGround, s.Vel.Y, newVel.Y)
	if s.onGround || (s.Vel.Y != newVel.Y && s.Vel.Y < 0) {
		bb := s.BB()
		//fmt.Printf("Player pos = %0.2f, %0.2f, %0.2f\n", bb.X.Min, bb.Y.Min, bb.Z.Min)
		surroundings := s.surroundings(bb.Offset(s.Vel.X, stepHeight, s.Vel.Z), w)
		outVel := s.Vel

		outVel.Y = stepHeight
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
		//fmt.Printf("Post-collision = %0.2f, %0.2f, %0.2f\n", bb.X.Min, bb.Y.Min, bb.Z.Min)

		outVel.Y *= -1
		// Lower the player back down to be on the ground.
		for _, b := range surroundings {
			outVel.Y = b.YOffset(bb, outVel.Y)
		}
		bb = bb.Offset(0, outVel.Y, 0)
		//fmt.Printf("Post-lower = %0.2f, %0.2f, %0.2f\n", bb.X.Min, bb.Y.Min, bb.Z.Min)

		oldMove := newVel.X*newVel.X + newVel.Z*newVel.Z
		newMove := outVel.X*outVel.X + outVel.Z*outVel.Z
		// fmt.Printf("oldMove/newmove = %v, (%0.2f >= %0.6f) = %v\n", oldMove >= newMove, outVel.Y, 0.000002-stepHeight, outVel.Y <= (0.000002-stepHeight))
		if oldMove >= newMove || outVel.Y <= (0.000002-stepHeight) {
			// fmt.Println("nope")
		} else {
			player = bb
			newVel = outVel
		}
	}

	// Update flags.
	s.Pos.X = player.X.Min + playerWidth/2
	s.Pos.Y = player.Y.Min
	s.Pos.Z = player.Z.Min + playerWidth/2
	s.collision.horizontal = newVel.X != s.Vel.X || newVel.Z != s.Vel.Z
	s.collision.vertical = newVel.Y != s.Vel.Y
	s.onGround = s.collision.vertical && s.Vel.Y < 0
	s.Vel = newVel
}

func modYaw(new, old float64) float64 {
	delta := math.Mod(new-old, 360)
	if delta > 180 {
		delta = 180 - delta
	} else if delta < -180 {
		delta += 360
	}
	// fmt.Printf("(%.2f - %.2f) = %.2f\n", new, old, delta)
	return delta
}

func (s *State) computeCollisionYXZ(bb, query AABB, vel path.Point, w World) (outBB AABB, outVel path.Point) {
	surroundings := s.surroundings(query, w)
	outVel = vel

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

// AtLookTarget returns true if the player look position is actually at the
// given pitch and yaw.
func (s *State) AtLookTarget(yaw, pitch float64) bool {
	dYaw, dPitch := math.Abs(modYaw(yaw, s.Yaw)), math.Abs(pitch-s.Pitch)
	return dYaw <= 0.8 && dPitch <= 1.1
}
