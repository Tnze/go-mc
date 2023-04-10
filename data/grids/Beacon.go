package grids

type Beacon struct {
	*Generic
	Type int
}

func NewBeacon() *Beacon {
	return &Beacon{
		Generic: InitGenericContainer(37, 1, 1),
		Type:    8,
	}
}
