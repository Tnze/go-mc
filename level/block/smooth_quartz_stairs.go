package block

type SmoothQuartzStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (SmoothQuartzStairs) ID() string {
	return "minecraft:smooth_quartz_stairs"
}
