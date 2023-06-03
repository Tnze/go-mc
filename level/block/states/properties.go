package states

import (
	"github.com/Tnze/go-mc/internal/utils"
	"golang.org/x/exp/constraints"
	"strconv"
)

type Property[Type any] interface {
	GetName() string
	GetValue(other string) any
	CanUpdate(other uint32) bool // other is either an int, PropertyEnum or bool
	GetValues() []any
	HashCode() uint64
}

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

func NewPropertyInteger(name string, min, max int) *PropertyInteger[int] {
	return &PropertyInteger[int]{
		name:   name,
		min:    min,
		max:    max,
		values: fillInteger(min, max),
	}
}

func (p PropertyInteger[Type]) GetName() string {
	return p.name
}

func (p PropertyInteger[Type]) GetValue(other string) any {
	if value, err := strconv.Atoi(other); err == nil {
		if value >= p.min && value <= p.max {
			return value
		}
	}
	return p.min
}

func (p PropertyInteger[Type]) CanUpdate(other uint32) bool {
	return p.min <= int(other) && int(other) <= p.max
}

func (p PropertyInteger[Type]) GetValues() []any {
	values := make([]any, len(p.values))
	for i, value := range p.values {
		values[i] = value
	}
	return values
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

func NewPropertyBoolean(name string) *PropertyBoolean[bool] {
	return &PropertyBoolean[bool]{
		name:   name,
		values: []bool{true, false},
	}
}

func (p PropertyBoolean[Type]) GetName() string {
	return p.name
}

func (p PropertyBoolean[Type]) GetValue(other string) any {
	if value, err := strconv.ParseBool(other); err == nil {
		return value
	}
	return false
}

func (p PropertyBoolean[Type]) CanUpdate(other uint32) bool {
	return true // For now we will assume that all booleans are valid
}

func (p PropertyBoolean[Type]) GetValues() []any {
	values := make([]any, len(p.values))
	for i, value := range p.values {
		values[i] = value
	}
	return values
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

func NewPropertyEnum[Type PropertiesEnum](name string, names map[string]Type) *PropertyEnum[Type] {
	return &PropertyEnum[Type]{
		name:  name,
		names: names,
	}
}

func (p PropertyEnum[Type]) GetName() string {
	return p.name
}

func (p PropertyEnum[Type]) GetValue(other string) any {
	for _, value := range p.names {
		if value.String() == other {
			return value
		}
	}
	panic("invalid value")
}

func (p PropertyEnum[Type]) CanUpdate(other uint32) bool {
	for _, value := range p.names {
		if value.Value() == byte(other) {
			return true
		}
	}
	return false
}

func (p PropertyEnum[Type]) GetValues() []any {
	var values []any
	for _, value := range p.names {
		values = append(values, value)
	}
	return values
}

func (p PropertyEnum[Type]) HashCode() uint64 {
	h := uint64(17) // choose a random initial value
	h = h*31 + utils.HashString(p.name)
	for _, value := range p.names {
		h = h*31 + uint64(len(value.String()))
	}
	return h
}
