package block

type BlueWallBanner struct {
	Facing string
}

func (BlueWallBanner) ID() string {
	return "minecraft:blue_wall_banner"
}
