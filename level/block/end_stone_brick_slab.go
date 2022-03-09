package block

type EndStoneBrickSlab struct {
	Type        string
	Waterlogged string
}

func (EndStoneBrickSlab) ID() string {
	return "minecraft:end_stone_brick_slab"
}
