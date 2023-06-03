package grids

type Generic9x6 struct {
	*Generic
}

func NewGeneric9x6(inventory *GenericInventory) *Generic9x6 {
	return &Generic9x6{InitGenericContainer("minecraft:generic_9x6", 5, 54, inventory)}
}
