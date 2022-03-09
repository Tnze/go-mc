package block

type YellowStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (YellowStainedGlassPane) ID() string {
	return "minecraft:yellow_stained_glass_pane"
}
