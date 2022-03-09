package block

type LightBlueBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (LightBlueBed) ID() string {
	return "minecraft:light_blue_bed"
}
