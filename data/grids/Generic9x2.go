package grids

type Generic9x2 struct {
	*Generic
}

func NewGeneric9x2(inventory *GenericInventory) *Generic9x2 {
	return &Generic9x2{InitGenericContainer("minecraft:generic_9x2", 1, 18, inventory)}
}
