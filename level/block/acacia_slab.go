package block

type AcaciaSlab struct {
	Type        string
	Waterlogged string
}

func (AcaciaSlab) ID() string {
	return "minecraft:acacia_slab"
}
