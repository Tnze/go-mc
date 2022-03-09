package block

type CobblestoneSlab struct {
	Type        string
	Waterlogged string
}

func (CobblestoneSlab) ID() string {
	return "minecraft:cobblestone_slab"
}
