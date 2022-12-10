package grids

type Generic_3x3 struct {
	*Generic
	Type int
}

func NewGeneric_3x3() *Generic_3x3 {
	return &Generic_3x3{
		Generic: InitGenericContainer(53, 3, 3),
		Type:    6,
	}
}
