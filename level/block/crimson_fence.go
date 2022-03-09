package block

type CrimsonFence struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (CrimsonFence) ID() string {
	return "minecraft:crimson_fence"
}
