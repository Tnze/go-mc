// Package path implements pathfinding.
package path

import (
	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/data/block"
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

// Movement represents a single type of movement in a path.
type Movement uint8

var allMovements = []Movement{TraverseNorth, TraverseSouth, TraverseEast, TraverseWest}

// Valid movement values.
const (
	Waypoint Movement = iota
	TraverseNorth
	TraverseSouth
	TraverseEast
	TraverseWest
)

// Tile represents a point in a path. All tiles in a path are adjaceent their
// preceeding tiles.
type Tile struct {
	Nav *Nav

	Movement Movement
	Pos      V3
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

	if t.Pos == t.Nav.Dest && t.Movement != Waypoint {
		dupe := t
		dupe.Movement = Waypoint
		return []astar.Pather{dupe}
	}

	for _, m := range allMovements {
		x, y, z := m.Offset()
		pos := V3{X: t.Pos.X + x, Y: t.Pos.Y + y, Z: t.Pos.Z + z}
		if m.Possible(t.Nav, pos.X, pos.Y, pos.Z) {
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

func (m Movement) Possible(nav *Nav, x, y, z int) bool {
	// fmt.Printf("%s.Possible(%d,%d,%d)\n", m, x, y, z)
	switch m {
	case Waypoint, TraverseNorth, TraverseSouth, TraverseEast, TraverseWest:
		b := nav.World.GetBlockStatus(x, y, z)
		if _, safe := safeStepBlocks[b]; !safe {
			return false
		}
		above1 := uint32(nav.World.GetBlockStatus(x, y+1, z))
		above2 := uint32(nav.World.GetBlockStatus(x, y+2, z))
		return above1 == block.Air.MinStateID && above2 == block.Air.MinStateID
	default:
		panic(m)
	}
}

func (m Movement) Offset() (x, y, z int) {
	switch m {
	case Waypoint:
		return 0, 0, 0
	case TraverseNorth:
		return 0, 0, -1
	case TraverseSouth:
		return 0, 0, 1
	case TraverseEast:
		return 1, 0, 0
	case TraverseWest:
		return -1, 0, 0
	default:
		panic(m)
	}
}

func (m Movement) BaseCost() float64 {
	switch m {
	case Waypoint:
		return 0
	case TraverseNorth, TraverseSouth, TraverseEast, TraverseWest:
		return 1
	default:
		panic(m)
	}
}

func (m Movement) String() string {
	switch m {
	case Waypoint:
		return "waypoint"
	case TraverseNorth:
		return "traverse-north"
	case TraverseSouth:
		return "traverse-south"
	case TraverseEast:
		return "traverse-east"
	case TraverseWest:
		return "traverse-west"
	default:
		panic(m)
	}
}
