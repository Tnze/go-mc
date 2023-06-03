package grids

type BrewingStand struct {
	*Generic
}

func NewBrewingStand(inventory *GenericInventory) *BrewingStand {
	return &BrewingStand{InitGenericContainer("minecraft:brewing_stand", 10, 5, inventory)}
}
