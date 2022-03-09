package block

type SandstoneSlab struct {
	Type        string
	Waterlogged string
}

func (SandstoneSlab) ID() string {
	return "minecraft:sandstone_slab"
}
