package block

type PrismarineSlab struct {
	Type        string
	Waterlogged string
}

func (PrismarineSlab) ID() string {
	return "minecraft:prismarine_slab"
}
