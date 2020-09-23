package path

// Movement represents a single type of movement in a path.
type Movement uint8

var allMovements = []Movement{TraverseNorth, TraverseSouth, TraverseEast, TraverseWest,
	TraverseNorthWest, TraverseNorthEast, TraverseSouthWest, TraverseSouthEast,
	DropNorth, DropSouth, DropEast, DropWest,
	AscendNorth, AscendSouth, AscendEast, AscendWest,
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
)

func (m Movement) Possible(nav *Nav, x, y, z int, from V3) bool {
	// fmt.Printf("%s.Possible(%d,%d,%d)\n", m, x, y, z)
	switch m {
	case Waypoint, TraverseNorth, TraverseSouth, TraverseEast, TraverseWest:
		if !SteppableBlock(nav.World.GetBlockStatus(x, y, z)) {
			return false
		}
		return AirLikeBlock(nav.World.GetBlockStatus(x, y+1, z)) && AirLikeBlock(nav.World.GetBlockStatus(x, y+2, z))

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

	case DropNorth:
		return 0, -1, -1
	case DropSouth:
		return 0, -1, 1
	case DropEast:
		return 1, -1, 0
	case DropWest:
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
	default:
		panic(m)
	}
}
