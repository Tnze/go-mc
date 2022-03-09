package block

type WaterCauldron struct {
	Level string
}

func (WaterCauldron) ID() string {
	return "minecraft:water_cauldron"
}
