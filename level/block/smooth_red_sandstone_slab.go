package block

type SmoothRedSandstoneSlab struct {
	Type        string
	Waterlogged string
}

func (SmoothRedSandstoneSlab) ID() string {
	return "minecraft:smooth_red_sandstone_slab"
}
