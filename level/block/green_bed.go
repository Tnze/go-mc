package block

type GreenBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (GreenBed) ID() string {
	return "minecraft:green_bed"
}
