package block

type QuartzPillar struct {
	Axis string
}

func (QuartzPillar) ID() string {
	return "minecraft:quartz_pillar"
}
