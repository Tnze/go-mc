package block

type DarkOakLeaves struct {
	Distance   string
	Persistent string
}

func (DarkOakLeaves) ID() string {
	return "minecraft:dark_oak_leaves"
}
