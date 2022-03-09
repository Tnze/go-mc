package block

type LimeStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (LimeStainedGlassPane) ID() string {
	return "minecraft:lime_stained_glass_pane"
}
