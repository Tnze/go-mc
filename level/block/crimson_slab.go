package block

type CrimsonSlab struct {
	Type        string
	Waterlogged string
}

func (CrimsonSlab) ID() string {
	return "minecraft:crimson_slab"
}
