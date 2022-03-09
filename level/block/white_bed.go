package block

type WhiteBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (WhiteBed) ID() string {
	return "minecraft:white_bed"
}
