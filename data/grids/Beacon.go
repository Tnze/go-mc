package grids

type Beacon struct {
	*Generic
}

func NewBeacon(inventory *GenericInventory) *Beacon {
	return &Beacon{InitGenericContainer("minecraft:beacon", 8, 1, inventory)}
}
