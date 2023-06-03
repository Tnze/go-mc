package grids

type Generic3x3 struct {
	*Generic
}

func NewGeneric3x3(inventory *GenericInventory) *Generic3x3 {
	return &Generic3x3{InitGenericContainer("minecraft:generic_3x3", 6, 9, inventory)}
}
