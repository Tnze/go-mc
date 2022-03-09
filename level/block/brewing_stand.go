package block

type BrewingStand struct {
	Has_bottle_0 string
	Has_bottle_1 string
	Has_bottle_2 string
}

func (BrewingStand) ID() string {
	return "minecraft:brewing_stand"
}
