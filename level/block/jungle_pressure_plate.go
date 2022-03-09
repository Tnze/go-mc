package block

type JunglePressurePlate struct {
	Powered string
}

func (JunglePressurePlate) ID() string {
	return "minecraft:jungle_pressure_plate"
}
