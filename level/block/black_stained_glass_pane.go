package block

type BlackStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (BlackStainedGlassPane) ID() string {
	return "minecraft:black_stained_glass_pane"
}
