package grids

type Generic_9x2 struct {
	*Generic
	Type int
}

func NewGeneric_9x2() *Generic_9x2 {
	return &Generic_9x2{
		Generic: InitGenericContainer(53, 9, 2),
		Type:    1,
	}
}
