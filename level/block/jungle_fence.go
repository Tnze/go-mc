package block

type JungleFence struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (JungleFence) ID() string {
	return "minecraft:jungle_fence"
}
