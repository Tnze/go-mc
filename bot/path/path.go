// Package path implements pathfinding.
package path

import (
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/data/block"
	"github.com/beefsack/go-astar"
)

// Point represents a point in 3D space.
type Point struct {
	X, Y, Z float64
}

type V3 struct {
	X, Y, Z int
}

func (v V3) Cost(other V3) float64 {
	x, y, z := v.X-other.X, v.Y-other.Y, v.Z-other.Z
	return math.Sqrt(float64(x*x+z*z)) + math.Sqrt(1.2*float64(y*y))
}

// Nav represents a navigation to a position.
type Nav struct {
	World       *world.World
	Start, Dest V3
}

func (n *Nav) Path() (path []astar.Pather, distance float64, found bool) {
	return astar.Path(
		Tile{ // Start point
			Nav:      n,
			Movement: Waypoint,
			Pos:      n.Start,
		},
		Tile{ // Destination point
			Nav:      n,
			Movement: Waypoint,
			Pos:      n.Dest,
		})
}

// Tile represents a point in a path. All tiles in a path are adjaceent their
// preceeding tiles.
type Tile struct {
	Nav *Nav

	Movement    Movement
	Pos         V3
	BlockStatus world.BlockStatus
	ExtraCost   int
}

func (t Tile) PathNeighborCost(to astar.Pather) float64 {
	other := to.(Tile)
	return 1 + other.Movement.BaseCost()
}

func (t Tile) PathEstimatedCost(to astar.Pather) float64 {
	other := to.(Tile)
	cost := t.Pos.Cost(other.Pos)

	return cost + other.Movement.BaseCost()
}

func (t Tile) PathNeighbors() []astar.Pather {
	possibles := make([]astar.Pather, 0, 8)

	if c := t.PathEstimatedCost(Tile{Pos: t.Nav.Start}); c > 8000 {
		return nil
	}

	if t.Pos == t.Nav.Dest && t.Movement != Waypoint {
		dupe := t
		dupe.Movement = Waypoint
		dupe.BlockStatus = 0
		return []astar.Pather{dupe}
	}

	for _, m := range allMovements {
		x, y, z := m.Offset()
		pos := V3{X: t.Pos.X + x, Y: t.Pos.Y + y, Z: t.Pos.Z + z}
		possible := m.Possible(t.Nav, pos.X, pos.Y, pos.Z, t.Pos, t.Movement)
		// fmt.Printf("%v-%v: Trying (%v) %v: possible=%v\n", t.Movement, t.Pos, pos, m, possible)
		if possible {
			possibles = append(possibles, Tile{
				Nav:         t.Nav,
				Movement:    m,
				Pos:         pos,
				BlockStatus: t.Nav.World.GetBlockStatus(pos.X, pos.Y, pos.Z),
			})
		}
	}

	// fmt.Printf("%v.Neighbours(): %+v\n", t.Pos, possibles)
	return possibles
}

func (t Tile) Inputs(pos, deltaPos, vel Point, runTime time.Duration) Inputs {
	// Sufficient for simple movements.
	at := math.Atan2(-deltaPos.X, -deltaPos.Z)
	mdX, _, mdZ := t.Movement.Offset()
	wantYaw := -math.Atan2(float64(mdX), float64(mdZ)) * 180 / math.Pi
	out := Inputs{
		ThrottleX: math.Sin(at),
		ThrottleZ: math.Cos(at),
		Yaw:       wantYaw,
	}
	if mdX == 0 && mdZ == 0 {
		out.Yaw = math.NaN()
	}
	if (rand.Int() % 14) == 0 {
		out.Pitch = float64((rand.Int() % 4) - 2)
	}

	switch t.Movement {
	case DescendLadder, DescendLadderEast, DescendLadderWest, DescendLadderNorth, DescendLadderSouth:
		// Deadzone the throttle to prevent an accidental ascend.
		if dist2 := math.Sqrt(deltaPos.X*deltaPos.X + deltaPos.Z*deltaPos.Z); dist2 < (0.22 * 0.22 * 2) {
			out.ThrottleX, out.ThrottleZ = 0, 0
		}

	case AscendLadder:
		dist2 := math.Sqrt(deltaPos.X*deltaPos.X + deltaPos.Z*deltaPos.Z)

		if x, _, z := LadderDirection(t.BlockStatus).Offset(); dist2 > (0.8*0.8) && deltaPos.Y < 0 {
			pos.X -= (0.25 * float64(x))
			pos.Z -= (0.25 * float64(z))
		} else {
			pos.X += (0.42 * float64(x))
			pos.Z += (0.42 * float64(z))
		}

		at = math.Atan2(-pos.X+float64(t.Pos.X)+0.5, -pos.Z+float64(t.Pos.Z)+0.5)
		out = Inputs{
			ThrottleX: math.Sin(at),
			ThrottleZ: math.Cos(at),
			Yaw:       math.NaN(),
		}

	case AscendNorth, AscendSouth, AscendEast, AscendWest:
		var (
			b          = block.ByID[block.StateID[uint32(t.BlockStatus)]]
			isStairs   = strings.HasSuffix(b.Name, "_stairs")
			maybeStuck = runTime < 1250*time.Millisecond
			dist2      = math.Sqrt(deltaPos.X*deltaPos.X + deltaPos.Z*deltaPos.Z)
		)
		out.Jump = dist2 < 1.75 && deltaPos.Y < -0.81

		// Special logic for stairs: Try to go towards the downwards edge initially.
		if isStairs && dist2 > (0.9*0.9) && deltaPos.Y < 0 {
			if x, _, z := StairsDirection(t.BlockStatus).Offset(); dist2 > (0.9*0.9) && deltaPos.Y < 0 {
				pos.X += (0.49 * float64(x))
				pos.Z += (0.49 * float64(z))
			}

			at = math.Atan2(-pos.X+float64(t.Pos.X)+0.5, -pos.Z+float64(t.Pos.Z)+0.5)
			out = Inputs{
				ThrottleX: math.Sin(at),
				ThrottleZ: math.Cos(at),
				Yaw:       math.NaN(),
				Jump:      out.Jump && !maybeStuck,
			}
		}

		// Turn off the throttle if we get stuck on the jump.
		if dist2 < 1 && deltaPos.Y < 0 && vel.Y == 0 {
			out.ThrottleX, out.ThrottleZ = 0, 0
		}

	case JumpCrossEast, JumpCrossWest, JumpCrossNorth, JumpCrossSouth:
		dist2 := math.Sqrt(deltaPos.X*deltaPos.X + deltaPos.Z*deltaPos.Z)
		out.Jump = dist2 > 1.5 && dist2 < 1.78
	}
	return out
}

func (t Tile) IsComplete(d Point) bool {
	switch t.Movement {
	case DescendLadder, DescendLadderNorth, DescendLadderSouth, DescendLadderWest, DescendLadderEast,
		DropNorth, DropSouth, DropEast, DropWest:
		return (d.X*d.X+d.Z*d.Z) < (2*0.2*0.25) && d.Y <= 0.05
	case AscendLadder:
		return d.Y >= 0
	case JumpCrossEast, JumpCrossWest, JumpCrossNorth, JumpCrossSouth:
		return (d.X*d.X+d.Z*d.Z) < (0.22*0.22) && d.Y >= -0.065
	}

	return (d.X*d.X+d.Z*d.Z) < (0.18*0.18) && d.Y >= -0.065 && d.Y <= 0.08
}
