package block

type BlackstoneSlab struct {
	Type        string
	Waterlogged string
}

func (BlackstoneSlab) ID() string {
	return "minecraft:blackstone_slab"
}
