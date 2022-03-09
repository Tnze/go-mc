package block

type PolishedBlackstoneBrickSlab struct {
	Type        string
	Waterlogged string
}

func (PolishedBlackstoneBrickSlab) ID() string {
	return "minecraft:polished_blackstone_brick_slab"
}
