package block

type BirchPressurePlate struct {
	Powered string
}

func (BirchPressurePlate) ID() string {
	return "minecraft:birch_pressure_plate"
}
