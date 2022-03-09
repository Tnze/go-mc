package block

type JungleSlab struct {
	Type        string
	Waterlogged string
}

func (JungleSlab) ID() string {
	return "minecraft:jungle_slab"
}
