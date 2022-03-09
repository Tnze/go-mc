package block

type OrangeStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (OrangeStainedGlassPane) ID() string {
	return "minecraft:orange_stained_glass_pane"
}
