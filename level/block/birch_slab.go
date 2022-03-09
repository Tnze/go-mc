package block

type BirchSlab struct {
	Type        string
	Waterlogged string
}

func (BirchSlab) ID() string {
	return "minecraft:birch_slab"
}
