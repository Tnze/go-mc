package grids

type Loom struct {
	*Generic
	Type int
}

func NewLoom() *Loom {
	return &Loom{
		Generic: InitGenericContainer(41, 4, 1),
		Type:    17,
	}
}
