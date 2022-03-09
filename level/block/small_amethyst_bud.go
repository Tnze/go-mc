package block

type SmallAmethystBud struct {
	Facing      string
	Waterlogged string
}

func (SmallAmethystBud) ID() string {
	return "minecraft:small_amethyst_bud"
}
