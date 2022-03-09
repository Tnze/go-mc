package block

type StonePressurePlate struct {
	Powered string
}

func (StonePressurePlate) ID() string {
	return "minecraft:stone_pressure_plate"
}
