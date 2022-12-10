package grids

type Anvil struct {
	*Generic
	Type int
}

func NewAnvil() *Anvil {
	return &Anvil{
		Generic: InitGenericContainer(39, 3, 1),
		Type:    7,
	}
}
