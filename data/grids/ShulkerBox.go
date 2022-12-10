package grids

type ShulkerBox struct {
	*Generic
	Type int
}

func NewShulkerBox() *ShulkerBox {
	return &ShulkerBox{
		Generic: InitGenericContainer(27, 3, 1),
		Type:    19,
	}
}
