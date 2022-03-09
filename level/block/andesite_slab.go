package block

type AndesiteSlab struct {
	Type        string
	Waterlogged string
}

func (AndesiteSlab) ID() string {
	return "minecraft:andesite_slab"
}
