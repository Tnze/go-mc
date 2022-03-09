package block

type CyanStainedGlassPane struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (CyanStainedGlassPane) ID() string {
	return "minecraft:cyan_stained_glass_pane"
}
