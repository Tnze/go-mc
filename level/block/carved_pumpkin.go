package block

type CarvedPumpkin struct {
	Facing string
}

func (CarvedPumpkin) ID() string {
	return "minecraft:carved_pumpkin"
}
