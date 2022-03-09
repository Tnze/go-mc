package block

type CrimsonPressurePlate struct {
	Powered string
}

func (CrimsonPressurePlate) ID() string {
	return "minecraft:crimson_pressure_plate"
}
