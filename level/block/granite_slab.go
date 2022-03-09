package block

type GraniteSlab struct {
	Type        string
	Waterlogged string
}

func (GraniteSlab) ID() string {
	return "minecraft:granite_slab"
}
