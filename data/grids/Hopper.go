package grids

type Hopper struct { // Also minecart with hopper
	*Generic
}

func NewHopper(inventory *GenericInventory) *Hopper {
	return &Hopper{InitGenericContainer("minecraft:hopper", 15, 5, inventory)}
}
