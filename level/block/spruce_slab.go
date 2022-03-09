package block

type SpruceSlab struct {
	Type        string
	Waterlogged string
}

func (SpruceSlab) ID() string {
	return "minecraft:spruce_slab"
}
