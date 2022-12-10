package grids

type BlastFurnace struct {
	*Generic
	Type int
}

func NewBlastFurnace() *BlastFurnace {
	return &BlastFurnace{
		Generic: InitGenericContainer(39, 3, 1),
		Type:    9,
	}
}
