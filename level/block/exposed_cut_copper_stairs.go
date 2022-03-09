package block

type ExposedCutCopperStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (ExposedCutCopperStairs) ID() string {
	return "minecraft:exposed_cut_copper_stairs"
}
