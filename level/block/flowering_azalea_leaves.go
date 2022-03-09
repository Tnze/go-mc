package block

type FloweringAzaleaLeaves struct {
	Distance   string
	Persistent string
}

func (FloweringAzaleaLeaves) ID() string {
	return "minecraft:flowering_azalea_leaves"
}
