package block

type GrayBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (GrayBed) ID() string {
	return "minecraft:gray_bed"
}
