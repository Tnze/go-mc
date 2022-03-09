package block

type OakPressurePlate struct {
	Powered string
}

func (OakPressurePlate) ID() string {
	return "minecraft:oak_pressure_plate"
}
