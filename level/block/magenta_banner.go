package block

type MagentaBanner struct {
	Rotation string
}

func (MagentaBanner) ID() string {
	return "minecraft:magenta_banner"
}
