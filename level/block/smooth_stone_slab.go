package block

type SmoothStoneSlab struct {
	Type        string
	Waterlogged string
}

func (SmoothStoneSlab) ID() string {
	return "minecraft:smooth_stone_slab"
}
