package block

type WarpedStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (WarpedStairs) ID() string {
	return "minecraft:warped_stairs"
}
