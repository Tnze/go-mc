package block

type PolishedAndesiteSlab struct {
	Type        string
	Waterlogged string
}

func (PolishedAndesiteSlab) ID() string {
	return "minecraft:polished_andesite_slab"
}
