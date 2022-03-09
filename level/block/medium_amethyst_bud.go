package block

type MediumAmethystBud struct {
	Facing      string
	Waterlogged string
}

func (MediumAmethystBud) ID() string {
	return "minecraft:medium_amethyst_bud"
}
