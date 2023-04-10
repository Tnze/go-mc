package grids

type Grindstone struct {
	*Generic
	Type int
}

func NewGrindstone() *Grindstone {
	return &Grindstone{
		Generic: InitGenericContainer(39, 3, 1),
		Type:    14,
	}
}
