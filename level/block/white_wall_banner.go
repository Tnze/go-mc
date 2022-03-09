package block

type WhiteWallBanner struct {
	Facing string
}

func (WhiteWallBanner) ID() string {
	return "minecraft:white_wall_banner"
}
