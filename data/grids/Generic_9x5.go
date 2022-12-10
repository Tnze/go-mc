package grids

type Generic_9x5 struct {
	*Generic
	Type int
}

func NewGeneric_9x5() *Generic_9x4 {
	return &Generic_9x4{
		Generic: InitGenericContainer(70, 9, 5),
		Type:    4,
	}
}
