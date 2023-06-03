package grids

type Generic9x1 struct {
	*Generic
}

func NewGeneric9x1(inventory *GenericInventory) *Generic9x1 {
	return &Generic9x1{InitGenericContainer("minecraft:generic_9x1", 0, 9, inventory)}
}
