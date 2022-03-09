package block

type MossyStoneBrickSlab struct {
	Type        string
	Waterlogged string
}

func (MossyStoneBrickSlab) ID() string {
	return "minecraft:mossy_stone_brick_slab"
}
