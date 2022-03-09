package block

type RedStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (RedStainedGlassPane) ID() string {
	return "minecraft:red_stained_glass_pane"
}
