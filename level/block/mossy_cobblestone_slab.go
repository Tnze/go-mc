package block

type MossyCobblestoneSlab struct {
	Type        string
	Waterlogged string
}

func (MossyCobblestoneSlab) ID() string {
	return "minecraft:mossy_cobblestone_slab"
}
