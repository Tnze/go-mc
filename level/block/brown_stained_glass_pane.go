package block

type BrownStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (BrownStainedGlassPane) ID() string {
	return "minecraft:brown_stained_glass_pane"
}
