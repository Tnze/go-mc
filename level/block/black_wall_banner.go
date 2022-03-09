package block

type BlackWallBanner struct {
	Facing string
}

func (BlackWallBanner) ID() string {
	return "minecraft:black_wall_banner"
}
