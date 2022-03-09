package block

type BrickSlab struct {
	Type        string
	Waterlogged string
}

func (BrickSlab) ID() string {
	return "minecraft:brick_slab"
}
