package block

type DioriteSlab struct {
	Type        string
	Waterlogged string
}

func (DioriteSlab) ID() string {
	return "minecraft:diorite_slab"
}
