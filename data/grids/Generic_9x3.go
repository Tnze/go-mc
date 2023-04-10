package grids

type Generic_9x3 struct {
	*Generic
	Type int
}

func NewGeneric_9x3() *Generic_9x3 {
	return &Generic_9x3{
		Generic: InitGenericContainer(62, 9, 3),
		Type:    2,
	}
}
