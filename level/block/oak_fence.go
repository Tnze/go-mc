package block

type OakFence struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (OakFence) ID() string {
	return "minecraft:oak_fence"
}
