package block

type SprucePressurePlate struct {
	Powered string
}

func (SprucePressurePlate) ID() string {
	return "minecraft:spruce_pressure_plate"
}
