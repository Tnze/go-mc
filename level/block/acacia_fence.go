package block

type AcaciaFence struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (AcaciaFence) ID() string {
	return "minecraft:acacia_fence"
}
