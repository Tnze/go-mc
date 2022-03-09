package block

type YellowWallBanner struct {
	Facing string
}

func (YellowWallBanner) ID() string {
	return "minecraft:yellow_wall_banner"
}
