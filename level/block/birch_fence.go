package block

type BirchFence struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (BirchFence) ID() string {
	return "minecraft:birch_fence"
}
