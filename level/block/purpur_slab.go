package block

type PurpurSlab struct {
	Type        string
	Waterlogged string
}

func (PurpurSlab) ID() string {
	return "minecraft:purpur_slab"
}
