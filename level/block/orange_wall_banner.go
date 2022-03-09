package block

type OrangeWallBanner struct {
	Facing string
}

func (OrangeWallBanner) ID() string {
	return "minecraft:orange_wall_banner"
}
