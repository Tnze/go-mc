package block

type QuartzSlab struct {
	Type        string
	Waterlogged string
}

func (QuartzSlab) ID() string {
	return "minecraft:quartz_slab"
}
