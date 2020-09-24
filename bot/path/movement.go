package path

import (
	"fmt"

	"github.com/Tnze/go-mc/bot/world"
	"github.com/Tnze/go-mc/data/block"
)

// Cardinal directions.
type Direction uint8

func (d Direction) Offset() (x, y, z int) {
	switch d {
	case North:
		return 0, 0, -1
	case South:
		return 0, 0, 1
	case East:
		return 1, 0, 0
	case West:
		return -1, 0, 0
	}
	panic(fmt.Sprintf("unknown direction value: %v", d))
}

func (d Direction) String() string {
	switch d {
	case North:
		return "north"
	case South:
		return "south"
	case East:
		return "east"
	case West:
		return "west"
	}
	return fmt.Sprintf("Direction?<%d>", int(d))
}

// Valid direction values.
const (
	North Direction = iota
	South
	West
	East
)

func LadderDirection(bStateID world.BlockStatus) Direction {
	return Direction(((uint32(bStateID) - block.Ladder.MinStateID) & 0xE) >> 1)
}

// Movement represents a single type of movement in a path.
type Movement uint8

var allMovements = []Movement{TraverseNorth, TraverseSouth, TraverseEast, TraverseWest,
	TraverseNorthWest, TraverseNorthEast, TraverseSouthWest, TraverseSouthEast,
	DropNorth, DropSouth, DropEast, DropWest,
	AscendNorth, AscendSouth, AscendEast, AscendWest,
	DescendLadder, DescendLadderNorth, DescendLadderSouth, DescendLadderEast, DescendLadderWest,
	AscendLadder,
}

// Valid movement values.
const (
	Waypoint Movement = iota
	TraverseNorth
	TraverseSouth
	TraverseEast
	TraverseWest
	TraverseNorthEast
	TraverseNorthWest
	TraverseSouthEast
	TraverseSouthWest
	DropNorth
	DropSouth
	DropEast
	DropWest
	AscendNorth
	AscendSouth
	AscendEast
	AscendWest
	DescendLadder
	DescendLadderNorth
	DescendLadderSouth
	DescendLadderEast
	DescendLadderWest
	AscendLadder
)

func (m Movement) Possible(nav *Nav, x, y, z int, from V3, previous Movement) bool {
	// fmt.Printf("%s.Possible(%d,%d,%d)\n", m, x, y, z)
	switch m {
	case Waypoint, TraverseNorth, TraverseSouth, TraverseEast, TraverseWest:
		if !SteppableBlock(nav.World.GetBlockStatus(x, y, z)) {
			return false
		}
		b1, b2 := nav.World.GetBlockStatus(x, y+1, z), nav.World.GetBlockStatus(x, y+2, z)
		u1, u2 := AirLikeBlock(b1) || IsLadder(b1), AirLikeBlock(b2) || IsLadder(b2)
		return u1 && u2

	case TraverseNorthWest, TraverseNorthEast, TraverseSouthWest, TraverseSouthEast:
		if !SteppableBlock(nav.World.GetBlockStatus(x, y, z)) {
			return false
		}
		if !AirLikeBlock(nav.World.GetBlockStatus(x, y+1, z)) || !AirLikeBlock(nav.World.GetBlockStatus(x, y+2, z)) {
			return false
		}
		if !AirLikeBlock(nav.World.GetBlockStatus(from.X, y+1, z)) || !AirLikeBlock(nav.World.GetBlockStatus(from.X, y+2, z)) {
			return false
		}
		return AirLikeBlock(nav.World.GetBlockStatus(x, y+1, from.Z)) && AirLikeBlock(nav.World.GetBlockStatus(x, y+2, from.Z))

	case DropNorth, DropSouth, DropEast, DropWest:
		for amt := 0; amt < 3; amt++ {
			if !AirLikeBlock(nav.World.GetBlockStatus(x, y+amt+1, z)) {
				return false
			}
		}
		return SteppableBlock(nav.World.GetBlockStatus(x, y, z))

	case AscendNorth, AscendSouth, AscendEast, AscendWest:
		if !AirLikeBlock(nav.World.GetBlockStatus(x, y+1, z)) || !AirLikeBlock(nav.World.GetBlockStatus(x, y+2, z)) {
			return false
		}
		return SteppableBlock(nav.World.GetBlockStatus(x, y, z)) &&
			AirLikeBlock(nav.World.GetBlockStatus(from.X, from.Y+1, from.Z)) &&
			AirLikeBlock(nav.World.GetBlockStatus(from.X, from.Y+2, from.Z))

	case DescendLadder, AscendLadder:
		if bID := nav.World.GetBlockStatus(x, y+1, z); !AirLikeBlock(bID) && !IsLadder(bID) {
			return false
		}
		return IsLadder(nav.World.GetBlockStatus(x, y, z))

	case DescendLadderNorth, DescendLadderSouth, DescendLadderEast, DescendLadderWest:
		for amt := 0; amt < 2; amt++ {
			if bID := nav.World.GetBlockStatus(x, y+amt+1, z); !AirLikeBlock(bID) && !IsLadder(bID) {
				return false
			}
		}
		return IsLadder(nav.World.GetBlockStatus(x, y, z))

	default:
		panic(m)
	}
}

func (m Movement) Offset() (x, y, z int) {
	switch m {
	case Waypoint:
		return 0, 0, 0
	case TraverseNorth:
		return North.Offset()
	case TraverseSouth:
		return South.Offset()
	case TraverseEast:
		return East.Offset()
	case TraverseWest:
		return West.Offset()

	case AscendLadder:
		return 0, 1, 0
	case DescendLadder:
		return 0, -1, 0
	case DropNorth, DescendLadderNorth:
		return 0, -1, -1
	case DropSouth, DescendLadderSouth:
		return 0, -1, 1
	case DropEast, DescendLadderEast:
		return 1, -1, 0
	case DropWest, DescendLadderWest:
		return -1, -1, 0
	case AscendNorth:
		return 0, 1, -1
	case AscendSouth:
		return 0, 1, 1
	case AscendEast:
		return 1, 1, 0
	case AscendWest:
		return -1, 1, 0

	case TraverseNorthWest:
		return -1, 0, -1
	case TraverseNorthEast:
		return 1, 0, -1
	case TraverseSouthWest:
		return -1, 0, 1
	case TraverseSouthEast:
		return 1, 0, 1

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
	case TraverseNorthWest, TraverseNorthEast, TraverseSouthWest, TraverseSouthEast:
		return 1.25

	case DropNorth, DropSouth, DropEast, DropWest:
		return 2
	case AscendNorth, AscendSouth, AscendEast, AscendWest:
		return 2.25
	case DescendLadderNorth, DescendLadderSouth, DescendLadderEast, DescendLadderWest:
		return 1.5
	case DescendLadder:
		return 1.2
	case AscendLadder:
		return 1.5
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

	case DropNorth:
		return "drop-north"
	case DropSouth:
		return "drop-south"
	case DropEast:
		return "drop-east"
	case DropWest:
		return "drop-west"

	case AscendNorth:
		return "jump-north"
	case AscendSouth:
		return "jump-south"
	case AscendEast:
		return "jump-east"
	case AscendWest:
		return "jump-west"

	case TraverseNorthWest:
		return "traverse-northwest"
	case TraverseNorthEast:
		return "traverse-northeast"
	case TraverseSouthWest:
		return "traverse-southwest"
	case TraverseSouthEast:
		return "traverse-southeast"

	case DescendLadder:
		return "descend-ladder"
	case DescendLadderNorth:
		return "descend-ladder-north"
	case DescendLadderSouth:
		return "descend-ladder-south"
	case DescendLadderEast:
		return "descend-ladder-east"
	case DescendLadderWest:
		return "descend-ladder-west"

	case AscendLadder:
		return "ascend-ladder"
	default:
		panic(m)
	}
}
