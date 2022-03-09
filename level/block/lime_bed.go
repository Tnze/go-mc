package block

type LimeBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (LimeBed) ID() string {
	return "minecraft:lime_bed"
}
