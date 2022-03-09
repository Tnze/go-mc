package block

type CutRedSandstoneSlab struct {
	Type        string
	Waterlogged string
}

func (CutRedSandstoneSlab) ID() string {
	return "minecraft:cut_red_sandstone_slab"
}
