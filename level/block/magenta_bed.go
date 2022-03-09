package block

type MagentaBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (MagentaBed) ID() string {
	return "minecraft:magenta_bed"
}
