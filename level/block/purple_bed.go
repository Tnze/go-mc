package block

type PurpleBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (PurpleBed) ID() string {
	return "minecraft:purple_bed"
}
