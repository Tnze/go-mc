package block

type LightGrayWallBanner struct {
	Facing string
}

func (LightGrayWallBanner) ID() string {
	return "minecraft:light_gray_wall_banner"
}
