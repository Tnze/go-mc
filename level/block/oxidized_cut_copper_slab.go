package block

type OxidizedCutCopperSlab struct {
	Type        string
	Waterlogged string
}

func (OxidizedCutCopperSlab) ID() string {
	return "minecraft:oxidized_cut_copper_slab"
}
