package block

type PrismarineBrickSlab struct {
	Type        string
	Waterlogged string
}

func (PrismarineBrickSlab) ID() string {
	return "minecraft:prismarine_brick_slab"
}
