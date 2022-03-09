package block

type WhiteStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (WhiteStainedGlassPane) ID() string {
	return "minecraft:white_stained_glass_pane"
}
