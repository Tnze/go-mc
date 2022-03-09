package block

type AcaciaPressurePlate struct {
	Powered string
}

func (AcaciaPressurePlate) ID() string {
	return "minecraft:acacia_pressure_plate"
}
