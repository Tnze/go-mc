package grids

type Generic9x5 struct {
	*Generic
}

func NewGeneric9x5(inventory *GenericInventory) *Generic9x4 {
	return &Generic9x4{InitGenericContainer("minecraft:generic_9x5", 4, 45, inventory)}
}
