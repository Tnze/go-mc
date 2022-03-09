package block

type LightGrayBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (LightGrayBed) ID() string {
	return "minecraft:light_gray_bed"
}
