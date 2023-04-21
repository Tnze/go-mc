package block

type StateDefinition[T any, K any] struct {
	States []K
}

func NewStateDefinition[T any, K any](states []K) StateDefinition[T, K] {
	return StateDefinition[T, K]{
		States: states,
	}
}

func (s StateDefinition[T, K]) Any() K {
	return s.States[0]
}
