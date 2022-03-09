package block

type PolishedBlackstoneSlab struct {
	Type        string
	Waterlogged string
}

func (PolishedBlackstoneSlab) ID() string {
	return "minecraft:polished_blackstone_slab"
}
