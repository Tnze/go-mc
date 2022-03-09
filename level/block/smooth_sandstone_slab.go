package block

type SmoothSandstoneSlab struct {
	Type        string
	Waterlogged string
}

func (SmoothSandstoneSlab) ID() string {
	return "minecraft:smooth_sandstone_slab"
}
