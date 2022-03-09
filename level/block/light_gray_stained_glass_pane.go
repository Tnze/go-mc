package block

type LightGrayStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (LightGrayStainedGlassPane) ID() string {
	return "minecraft:light_gray_stained_glass_pane"
}
