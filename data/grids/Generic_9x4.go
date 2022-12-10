package grids

type Generic_9x4 struct {
	*Generic
	Type int
}

func NewGeneric_9x4() *Generic_9x4 {
	return &Generic_9x4{
		Generic: InitGenericContainer(70, 9, 4),
		Type:    3,
	}
}
