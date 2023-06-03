package grids

type Grindstone struct {
	*Generic
}

func NewGrindstone(inventory *GenericInventory) *Grindstone {
	return &Grindstone{InitGenericContainer("minecraft:grindstone", 14, 3, inventory)}
}
