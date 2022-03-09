package block

type SpruceLeaves struct {
	Distance   string
	Persistent string
}

func (SpruceLeaves) ID() string {
	return "minecraft:spruce_leaves"
}
