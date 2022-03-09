package block

type BlueStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (BlueStainedGlassPane) ID() string {
	return "minecraft:blue_stained_glass_pane"
}
