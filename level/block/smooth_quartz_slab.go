package block

type SmoothQuartzSlab struct {
	Type        string
	Waterlogged string
}

func (SmoothQuartzSlab) ID() string {
	return "minecraft:smooth_quartz_slab"
}
