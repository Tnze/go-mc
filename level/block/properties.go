package block

import (
	"golang.org/x/exp/constraints"
	"strconv"
)

type Property[Type any] interface {
	GetValue(other string) Type
}

type Integer int
type Boolean bool

type PropertyInteger[Type constraints.Integer] struct {
	name     string
	min, max int
	values   []Type
}

func fillInteger(min, max int) []int {
	var values []int
	for i := min; i <= max; i++ {
		values = append(values, i)
	}
	return values
}

func NewPropertyInteger(name string, min, max int) PropertyInteger[int] {
	return PropertyInteger[int]{
		name:   name,
		min:    min,
		max:    max,
		values: fillInteger(min, max),
	}
}

func (p PropertyInteger[Type]) GetValue(other string) int {
	if value, err := strconv.Atoi(other); err == nil {
		return value
	}
	return p.min
}

type PropertyBoolean[Type bool] struct {
	name   string
	values []Type
}

func NewPropertyBoolean(name string) PropertyBoolean[bool] {
	return PropertyBoolean[bool]{
		name:   name,
		values: []bool{true, false},
	}
}

func (p PropertyBoolean[Type]) GetValue(other string) bool {
	if value, err := strconv.ParseBool(other); err == nil {
		return value
	}
	return false
}

type PropertyEnum[Type PropertiesEnum] struct {
	name  string
	names map[string]Type
}

func NewPropertyEnum[Type PropertiesEnum](name string, names map[string]Type) PropertyEnum[Type] {
	return PropertyEnum[Type]{
		name:  name,
		names: names,
	}
}

func (p PropertyEnum[Type]) GetValue(other string) Type {
	if value, ok := p.names[other]; ok {
		return value
	}
	panic("invalid value")
}

func (b Boolean) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatBool(bool(b))), nil
}

func (b *Boolean) UnmarshalText(text []byte) (err error) {
	*((*bool)(b)), err = strconv.ParseBool(string(text))
	return
}

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
