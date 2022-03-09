package block

type ExposedCutCopperSlab struct {
	Type        string
	Waterlogged string
}

func (ExposedCutCopperSlab) ID() string {
	return "minecraft:exposed_cut_copper_slab"
}
