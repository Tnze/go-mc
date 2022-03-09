package block

type RedWallBanner struct {
	Facing string
}

func (RedWallBanner) ID() string {
	return "minecraft:red_wall_banner"
}
