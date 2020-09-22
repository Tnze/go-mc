// Package path implements pathfinding.
package path

import (
	"math"

	"github.com/Tnze/go-mc/bot/world"
	"github.com/beefsack/go-astar"
)

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

	if t.PathEstimatedCost(Tile{Pos: t.Nav.Start}) > 1200 {
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
		if m.Possible(t.Nav, pos.X, pos.Y, pos.Z, t.Pos) {
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

func (t Tile) Inputs(dX, dY, dZ float64) Inputs {
	// Sufficient for simple movements.
	at := math.Atan2(-dX, -dZ)
	out := Inputs{
		ThrottleX: math.Sin(at),
		ThrottleZ: math.Cos(at),
	}

	switch t.Movement {
	case AscendNorth, AscendSouth, AscendEast, AscendWest:
		out.Jump = math.Sqrt(dX*dX+dZ*dZ) < 1.75
		out.Yaw = 0
	}
	return out
}
