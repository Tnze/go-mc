package block

type QuartzStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (QuartzStairs) ID() string {
	return "minecraft:quartz_stairs"
}
