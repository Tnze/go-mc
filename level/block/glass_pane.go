package block

type GlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (GlassPane) ID() string {
	return "minecraft:glass_pane"
}
