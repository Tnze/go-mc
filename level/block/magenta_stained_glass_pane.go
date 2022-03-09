package block

type MagentaStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (MagentaStainedGlassPane) ID() string {
	return "minecraft:magenta_stained_glass_pane"
}
