package block

type HeavyWeightedPressurePlate struct {
	Power string
}

func (HeavyWeightedPressurePlate) ID() string {
	return "minecraft:heavy_weighted_pressure_plate"
}
