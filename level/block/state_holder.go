package block

import (
	"github.com/Tnze/go-mc/internal/data"
	"github.com/Tnze/go-mc/level/block/states"
)

type StateHolder struct {
	Properties map[states.Property[any]]uint32
	Neighbors  data.HashTable[uint64, uint32]
}

func NewStateHolder(properties map[states.Property[any]]uint32) *StateHolder {
	return &StateHolder{
		Properties: properties,
		Neighbors:  *data.NewHashTable[uint64, uint32](),
	}
}

func (s *StateHolder) GetDefaultValue() StateID {
	return StateID(s.Neighbors.Get(0, 0))
}

func (s *StateHolder) GetValue(property states.Property[any], value uint32) StateID {
	return StateID(s.Neighbors.Get(property.HashCode(), value))
}

func (s *StateHolder) SetValue(property states.Property[any], value uint32, id StateID) any {
	s.Properties[property] = value
	var hashcode uint64
	if property != nil {
		hashcode = property.HashCode()
	} else {
		hashcode = 0 // nil property is like the address of the value 0 (zero)
	}
	s.Neighbors.Put(hashcode, value, int(id))
	return value
}

func (s *StateHolder) HasProperty(property states.Property[any]) bool {
	_, ok := s.Properties[property]
	return ok
}
