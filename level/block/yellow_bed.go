package block

type YellowBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (YellowBed) ID() string {
	return "minecraft:yellow_bed"
}
