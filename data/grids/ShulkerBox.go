package grids

type ShulkerBox struct {
	*Generic
}

func NewShulkerBox(inventory *GenericInventory) *ShulkerBox {
	return &ShulkerBox{InitGenericContainer("minecraft:shulker_box", 19, 27, inventory)}
}
