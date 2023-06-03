package block

import (
	"fmt"
	"github.com/Tnze/go-mc/level/block/states"
)

type StateFeeder[T any] struct {
	index int
	max   int
}

func NewStateFeeder[T any]() *StateFeeder[T] {
	return &StateFeeder[T]{0, 0}
}

func (s *StateFeeder[T]) FeedState(state *StateHolder, properties []states.Property[any]) *StateHolder {
	length := 1

	for _, p := range properties {
		length *= len(p.GetValues())
	}

	s.max += length
	if s.index > s.max {
		panic(fmt.Errorf("invalid state for length %v", length))
	}

	state.SetValue(nil, 0, StateID(s.index)) // default state

	if len(properties) == 0 {
		s.index++
		return state
	}

	for _, p := range properties {
		for _, v := range p.GetValues() {
			state.SetValue(p, parseState(v), StateID(s.index))
			s.index++
		}
	}

	return state
}
