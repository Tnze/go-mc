package grids

type Generic_9x6 struct {
	*Generic
	Type int
}

func NewGeneric_9x6() *Generic_9x6 {
	return &Generic_9x6{
		Generic: InitGenericContainer(78, 9, 6),
		Type:    5,
	}
}
