package block

type WarpedFence struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (WarpedFence) ID() string {
	return "minecraft:warped_fence"
}
