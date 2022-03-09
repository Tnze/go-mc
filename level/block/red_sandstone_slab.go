package block

type RedSandstoneSlab struct {
	Type        string
	Waterlogged string
}

func (RedSandstoneSlab) ID() string {
	return "minecraft:red_sandstone_slab"
}
