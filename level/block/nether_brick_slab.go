package block

type NetherBrickSlab struct {
	Type        string
	Waterlogged string
}

func (NetherBrickSlab) ID() string {
	return "minecraft:nether_brick_slab"
}
