package block

type GreenStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (GreenStainedGlassPane) ID() string {
	return "minecraft:green_stained_glass_pane"
}
