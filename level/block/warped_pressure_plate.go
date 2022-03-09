package block

type WarpedPressurePlate struct {
	Powered string
}

func (WarpedPressurePlate) ID() string {
	return "minecraft:warped_pressure_plate"
}
