package block

type PinkBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (PinkBed) ID() string {
	return "minecraft:pink_bed"
}
