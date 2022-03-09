package block

type AcaciaLeaves struct {
	Distance   string
	Persistent string
}

func (AcaciaLeaves) ID() string {
	return "minecraft:acacia_leaves"
}
