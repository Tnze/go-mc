package block

type LightWeightedPressurePlate struct {
	Power string
}

func (LightWeightedPressurePlate) ID() string {
	return "minecraft:light_weighted_pressure_plate"
}
