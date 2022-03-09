package block

type CyanBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (CyanBed) ID() string {
	return "minecraft:cyan_bed"
}
