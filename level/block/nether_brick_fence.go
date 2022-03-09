package block

type NetherBrickFence struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (NetherBrickFence) ID() string {
	return "minecraft:nether_brick_fence"
}
