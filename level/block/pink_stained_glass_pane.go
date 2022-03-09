package block

type PinkStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (PinkStainedGlassPane) ID() string {
	return "minecraft:pink_stained_glass_pane"
}
