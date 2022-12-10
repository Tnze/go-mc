package grids

type Stonercutter struct {
	*Generic
	Type int
}

func NewStonercutter() *Stonercutter {
	return &Stonercutter{
		Generic: InitGenericContainer(39, 2, 1),
		Type:    23,
	}
}
