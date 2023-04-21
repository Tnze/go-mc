package block

import "fmt"

type BlockState StateHolder[Block]

func NewBlockState(owner *Block, properties map[uint64]any) BlockState {
	return BlockState(NewStateHolder[Block](owner, properties))
}

func (s BlockState) GetValue(property Property[any]) any {
	return s.Properties[property.HashCode()]
}

func (s BlockState) SetValue(property Property[any], value any) {
	if !property.CanUpdate(value) {
		panic(fmt.Errorf("invalid value %v for property %v", value, property))
	}
	s.Properties[property.HashCode()] = value
}

func (s BlockState) GetOwner() *Block {
	return s.Owner
}

func (s BlockState) HasProperty(property Property[any]) bool {
	_, ok := s.Properties[property.HashCode()]
	return ok
}
