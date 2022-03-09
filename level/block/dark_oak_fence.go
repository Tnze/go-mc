package block

type DarkOakFence struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (DarkOakFence) ID() string {
	return "minecraft:dark_oak_fence"
}
