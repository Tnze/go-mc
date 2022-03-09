package block

type BlackBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (BlackBed) ID() string {
	return "minecraft:black_bed"
}
