package block

type DeepslateBrickSlab struct {
	Type        string
	Waterlogged string
}

func (DeepslateBrickSlab) ID() string {
	return "minecraft:deepslate_brick_slab"
}
