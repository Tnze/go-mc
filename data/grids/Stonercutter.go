package grids

type Stonecutter struct {
	*Generic
}

func NewStonecutter(inventory *GenericInventory) *Stonecutter {
	return &Stonecutter{InitGenericContainer("minecraft:stonecutter", 23, 2, inventory)}
}
