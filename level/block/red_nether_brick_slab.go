package block

type RedNetherBrickSlab struct {
	Type        string
	Waterlogged string
}

func (RedNetherBrickSlab) ID() string {
	return "minecraft:red_nether_brick_slab"
}
