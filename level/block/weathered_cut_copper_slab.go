package block

type WeatheredCutCopperSlab struct {
	Type        string
	Waterlogged string
}

func (WeatheredCutCopperSlab) ID() string {
	return "minecraft:weathered_cut_copper_slab"
}
