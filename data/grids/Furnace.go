package grids

type Furnace struct {
	*Generic
}

func NewFurnace(inventory *GenericInventory) *Furnace {
	return &Furnace{InitGenericContainer("minecraft:furnace", 13, 3, inventory)}
}
