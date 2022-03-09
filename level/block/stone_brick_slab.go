package block

type StoneBrickSlab struct {
	Type        string
	Waterlogged string
}

func (StoneBrickSlab) ID() string {
	return "minecraft:stone_brick_slab"
}
