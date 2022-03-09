package block

type WarpedSlab struct {
	Type        string
	Waterlogged string
}

func (WarpedSlab) ID() string {
	return "minecraft:warped_slab"
}
