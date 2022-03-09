package block

type GrayWallBanner struct {
	Facing string
}

func (GrayWallBanner) ID() string {
	return "minecraft:gray_wall_banner"
}
