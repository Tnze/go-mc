package block

type WaxedCutCopperSlab struct {
	Type        string
	Waterlogged string
}

func (WaxedCutCopperSlab) ID() string {
	return "minecraft:waxed_cut_copper_slab"
}
