package block

type MagentaWallBanner struct {
	Facing string
}

func (MagentaWallBanner) ID() string {
	return "minecraft:magenta_wall_banner"
}
