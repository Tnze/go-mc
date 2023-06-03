package grids

type Generic9x4 struct {
	*Generic
}

func NewGeneric9x4(inventory *GenericInventory) *Generic9x4 {
	return &Generic9x4{InitGenericContainer("minecraft:generic_9x4", 3, 36, inventory)}
}
