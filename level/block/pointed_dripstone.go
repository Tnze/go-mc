package block

type PointedDripstone struct {
	Thickness          string
	Vertical_direction string
	Waterlogged        string
}

func (PointedDripstone) ID() string {
	return "minecraft:pointed_dripstone"
}
