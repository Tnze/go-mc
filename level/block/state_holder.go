package block

type StateHolder[T any, K any] struct {
	Owner      T
	Properties map[Property[any]]K
}

func NewStateHolder[T any, K any](owner T, properties map[Property[any]]K) StateHolder[T, K] {
	return StateHolder[T, K]{
		Owner:      owner,
		Properties: properties,
	}
}

func (s StateHolder[T, K]) GetValue(property Property[any]) any {
	return s.Properties[property]
}

func (s StateHolder[T, K]) HasProperty(property Property[any]) bool {
	_, ok := s.Properties[property]
	return ok
}
