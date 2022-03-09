package block

type GrayStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (GrayStainedGlassPane) ID() string {
	return "minecraft:gray_stained_glass_pane"
}
