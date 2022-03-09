package block

type BrownBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (BrownBed) ID() string {
	return "minecraft:brown_bed"
}
