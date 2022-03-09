package block

type OakLeaves struct {
	Distance   string
	Persistent string
}

func (OakLeaves) ID() string {
	return "minecraft:oak_leaves"
}
