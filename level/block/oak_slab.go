package block

type OakSlab struct {
	Type        string
	Waterlogged string
}

func (OakSlab) ID() string {
	return "minecraft:oak_slab"
}
