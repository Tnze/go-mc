package block

type OrangeBed struct {
	Facing   string
	Occupied string
	Part     string
}

func (OrangeBed) ID() string {
	return "minecraft:orange_bed"
}
