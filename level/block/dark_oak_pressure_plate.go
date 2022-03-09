package block

type DarkOakPressurePlate struct {
	Powered string
}

func (DarkOakPressurePlate) ID() string {
	return "minecraft:dark_oak_pressure_plate"
}
