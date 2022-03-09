package block

type SculkSensor struct {
	Power              string
	Sculk_sensor_phase string
	Waterlogged        string
}

func (SculkSensor) ID() string {
	return "minecraft:sculk_sensor"
}
