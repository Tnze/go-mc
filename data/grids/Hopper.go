package grids

type Hopper struct { // Also minecart with hopper
	*Generic
	Type int
}

func NewHopper() *Hopper {
	return &Hopper{
		Generic: InitGenericContainer(41, 5, 1),
		Type:    15,
	}
}
