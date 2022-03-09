package block

type RedBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (RedBed) ID() string {
	return "minecraft:red_bed"
}
