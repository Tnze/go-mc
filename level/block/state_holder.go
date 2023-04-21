package block

type StateHolder[T any] struct {
	Owner      *T
	Properties map[uint64]any
}

func NewStateHolder[T any](owner *T, properties map[uint64]any) StateHolder[T] {
	return StateHolder[T]{
		Owner:      owner,
		Properties: properties,
	}
}

func (s StateHolder[T]) GetValue(property Property[any]) any {
	return s.Properties[property.HashCode()]
}

func (s StateHolder[T]) SetValue(property Property[any], value any) {
	s.Properties[property.HashCode()] = value
}

func (s StateHolder[T]) GetOwner() *T {
	return s.Owner
}

func (s StateHolder[T]) HasProperty(property Property[any]) bool {
	_, ok := s.Properties[property.HashCode()]
	return ok
}
