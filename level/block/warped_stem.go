package block

type WarpedStem struct {
	Axis string
}

func (WarpedStem) ID() string {
	return "minecraft:warped_stem"
}
