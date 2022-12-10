package grids

type Generic_9x1 struct {
	*Generic
	Type int
}

func NewGeneric_9x1() *Generic_9x1 {
	return &Generic_9x1{
		Generic: InitGenericContainer(44, 9, 1),
		Type:    0,
	}
}
