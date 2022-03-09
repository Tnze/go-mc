package block

type PurpleStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (PurpleStainedGlassPane) ID() string {
	return "minecraft:purple_stained_glass_pane"
}
