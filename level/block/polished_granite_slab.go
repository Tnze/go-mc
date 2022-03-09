package block

type PolishedGraniteSlab struct {
	Type        string
	Waterlogged string
}

func (PolishedGraniteSlab) ID() string {
	return "minecraft:polished_granite_slab"
}
