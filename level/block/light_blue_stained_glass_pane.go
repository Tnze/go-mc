package block

type LightBlueStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (LightBlueStainedGlassPane) ID() string {
	return "minecraft:light_blue_stained_glass_pane"
}
