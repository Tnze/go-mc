package block

type PetrifiedOakSlab struct {
	Type        string
	Waterlogged string
}

func (PetrifiedOakSlab) ID() string {
	return "minecraft:petrified_oak_slab"
}
