package block

type CutSandstoneSlab struct {
	Type        string
	Waterlogged string
}

func (CutSandstoneSlab) ID() string {
	return "minecraft:cut_sandstone_slab"
}
