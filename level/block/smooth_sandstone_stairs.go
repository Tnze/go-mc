package block

type SmoothSandstoneStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (SmoothSandstoneStairs) ID() string {
	return "minecraft:smooth_sandstone_stairs"
}
