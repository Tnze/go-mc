package block

type PolishedDioriteSlab struct {
	Type        string
	Waterlogged string
}

func (PolishedDioriteSlab) ID() string {
	return "minecraft:polished_diorite_slab"
}
