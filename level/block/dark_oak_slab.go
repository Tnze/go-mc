package block

type DarkOakSlab struct {
	Type        string
	Waterlogged string
}

func (DarkOakSlab) ID() string {
	return "minecraft:dark_oak_slab"
}
