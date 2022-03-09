package block

type BlueBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (BlueBed) ID() string {
	return "minecraft:blue_bed"
}
