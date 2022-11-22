package block

import "strconv"

//go:generate go run ./generator/properties/main.go

type Boolean bool

func (b Boolean) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatBool(bool(b))), nil
}

func (b *Boolean) UnmarshalText(text []byte) (err error) {
	*((*bool)(b)), err = strconv.ParseBool(string(text))
	return
}

type Integer int

func (i Integer) MarshalText() (text []byte, err error) {
	return []byte(strconv.Itoa(int(i))), nil
}

func (i *Integer) UnmarshalText(text []byte) (err error) {
	*((*int)(i)), err = strconv.Atoi(string(text))
	return
}

func (f FrontAndTop) Directions() (front, top Direction) {
	switch f {
	case DownEast:
		return Down, East
	case DownNorth:
		return Down, North
	case DownSouth:
		return Down, South
	case DownWest:
		return Down, West
	case UpEast:
		return Up, East
	case UpNorth:
		return Up, North
	case UpSouth:
		return Up, South
	case UpWest:
		return Up, West
	case WestUp:
		return West, Up
	case EastUp:
		return East, Up
	case NorthUp:
		return North, Up
	case SouthUp:
		return South, Up
	default:
		panic("invalid FrontAndTop")
	}
}
