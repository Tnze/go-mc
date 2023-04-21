package block

import (
	"github.com/Tnze/go-mc/internal/utils"
	"golang.org/x/exp/constraints"
	"strconv"
)

type Property[Type any] interface {
	GetValue(other string) any
	CanUpdate(other any) bool
	HashCode() uint64
}

type Integer int
type Boolean bool

type PropertyInteger[Type constraints.Integer] struct {
	Property[int]
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

func (p PropertyInteger[Type]) GetValue(other string) any {
	if value, err := strconv.Atoi(other); err == nil {
		if value >= p.min && value <= p.max {
			return value
		}
	}
	return p.min
}

func (p PropertyInteger[Type]) CanUpdate(other any) bool {
	if value, ok := other.(int); ok {
		return value >= p.min && value <= p.max
	}
	return false
}

func (p PropertyInteger[Type]) HashCode() uint64 {
	h := uint64(17) // choose a random initial value
	h = h*31 + utils.HashString(p.name)
	h = h*31 + uint64(p.min)
	h = h*31 + uint64(p.max)
	for _, value := range p.values {
		h = h*31 + uint64(value)
	}
	return h
}

type PropertyBoolean[Type bool] struct {
	Property[bool]
	name   string
	values []Type
}

func NewPropertyBoolean(name string) PropertyBoolean[bool] {
	return PropertyBoolean[bool]{
		name:   name,
		values: []bool{true, false},
	}
}

func (p PropertyBoolean[Type]) GetValue(other string) any {
	if value, err := strconv.ParseBool(other); err == nil {
		return value
	}
	return false
}

func (p PropertyBoolean[Type]) CanUpdate(other any) bool {
	if _, ok := other.(bool); ok {
		return true
	}
	return false
}

func (p PropertyBoolean[Type]) HashCode() uint64 {
	h := uint64(17) // choose a random initial value
	h = h*31 + utils.HashString(p.name)
	for _, value := range p.values {
		// Get the value from the binary
		var x int
		if value {
			x = 1
		} else {
			x = 0
		}
		h = h*31 + uint64(x)
	}
	return h
}

type PropertyEnum[Type PropertiesEnum] struct {
	Property[Type]
	name  string
	names map[string]Type
}

func NewPropertyEnum[Type PropertiesEnum](name string, names map[string]Type) PropertyEnum[Type] {
	return PropertyEnum[Type]{
		name:  name,
		names: names,
	}
}

func (p PropertyEnum[Type]) GetValue(other string) any {
	if value, ok := p.names[other]; ok {
		return value
	}
	panic("invalid value")
}

func (p PropertyEnum[Type]) CanUpdate(other any) bool {
	if _, ok := other.(PropertiesEnum); ok {
		return true
	}
	return false
}

func (p PropertyEnum[Type]) HashCode() uint64 {
	h := uint64(17) // choose a random initial value
	h = h*31 + utils.HashString(p.name)
	for _, value := range p.names {
		h = h*31 + uint64(len(value.String()))
	}
	return h
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
