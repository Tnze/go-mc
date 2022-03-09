package block

type CutCopperStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (CutCopperStairs) ID() string {
	return "minecraft:cut_copper_stairs"
}
