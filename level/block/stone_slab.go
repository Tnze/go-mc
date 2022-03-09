package block

type StoneSlab struct {
	Type        string
	Waterlogged string
}

func (StoneSlab) ID() string {
	return "minecraft:stone_slab"
}
