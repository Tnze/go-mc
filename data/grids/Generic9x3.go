package grids

type Generic9x3 struct {
	*Generic
}

func NewGeneric9x3(inventory *GenericInventory) *Generic9x3 {
	return &Generic9x3{InitGenericContainer("minecraft:generic_9x3", 2, 27, inventory)}
}
