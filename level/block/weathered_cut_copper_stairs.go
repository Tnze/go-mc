package block

type WeatheredCutCopperStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (WeatheredCutCopperStairs) ID() string {
	return "minecraft:weathered_cut_copper_stairs"
}
