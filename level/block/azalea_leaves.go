package block

type AzaleaLeaves struct {
	Distance   string
	Persistent string
}

func (AzaleaLeaves) ID() string {
	return "minecraft:azalea_leaves"
}
