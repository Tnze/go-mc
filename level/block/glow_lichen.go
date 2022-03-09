package block

type GlowLichen struct {
	Down        string
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (GlowLichen) ID() string {
	return "minecraft:glow_lichen"
}
