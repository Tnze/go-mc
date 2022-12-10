package grids

type Furnace struct {
	*Generic
	Type int
}

func NewFurnace() *Furnace {
	return &Furnace{
		Generic: InitGenericContainer(39, 3, 1),
		Type:    13,
	}
}
