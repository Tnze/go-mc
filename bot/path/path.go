// Package path implements pathfinding.
package path

import (
	"math"

	"github.com/Tnze/go-mc/bot/world"
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
	return float64(x*x+z*z) + 1.2*float64(y*y)
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

	Movement  Movement
	Pos       V3
	ExtraCost int
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
		return []astar.Pather{dupe}
	}

	for _, m := range allMovements {
		x, y, z := m.Offset()
		pos := V3{X: t.Pos.X + x, Y: t.Pos.Y + y, Z: t.Pos.Z + z}
		possible := m.Possible(t.Nav, pos.X, pos.Y, pos.Z, t.Pos, t.Movement)
		// fmt.Printf("%v-%v: Trying (%v) %v: possible=%v\n", t.Movement, t.Pos, pos, m, possible)
		if possible {
			possibles = append(possibles, Tile{
				Nav:      t.Nav,
				Movement: m,
				Pos:      pos,
			})
		}
	}

	// fmt.Printf("%v.Neighbours(): %+v\n", t.Pos, possibles)
	return possibles
}

func (t Tile) Inputs(deltaPos, vel Point) Inputs {
	// Sufficient for simple movements.
	at := math.Atan2(-deltaPos.X, -deltaPos.Z)
	out := Inputs{
		ThrottleX: math.Sin(at),
		ThrottleZ: math.Cos(at),
	}

	switch t.Movement {
	case AscendNorth, AscendSouth, AscendEast, AscendWest:
		dist2 := math.Sqrt(deltaPos.X*deltaPos.X + deltaPos.Z*deltaPos.Z)
		out.Jump = dist2 < 1.75 && deltaPos.Y < -0.81

		// Turn off the throttle if we get stuck on the jump.
		if dist2 < 1 && deltaPos.Y < 0 && vel.Y == 0 {
			out.ThrottleX, out.ThrottleZ = 0, 0
		}
		out.Yaw = 0
	}
	return out
}

func (t Tile) IsComplete(d Point) bool {
	switch t.Movement {
	case DescendLadder, DescendLadderNorth, DescendLadderSouth, DescendLadderWest, DescendLadderEast, DropNorth, DropSouth, DropEast, DropWest:
		return (d.X*d.X+d.Z*d.Z) < (0.1*0.1*0.13) && d.Y <= 0.05
	}

	return (d.X*d.X+d.Z*d.Z) < (0.18*0.18) && d.Y >= -0.01 && d.Y <= 0.08
}
