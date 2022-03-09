package block

type CutCopperSlab struct {
	Type        string
	Waterlogged string
}

func (CutCopperSlab) ID() string {
	return "minecraft:cut_copper_slab"
}
