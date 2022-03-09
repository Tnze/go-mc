package block

type DarkPrismarineSlab struct {
	Type        string
	Waterlogged string
}

func (DarkPrismarineSlab) ID() string {
	return "minecraft:dark_prismarine_slab"
}
