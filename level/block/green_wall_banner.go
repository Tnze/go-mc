package block

type GreenWallBanner struct {
	Facing string
}

func (GreenWallBanner) ID() string {
	return "minecraft:green_wall_banner"
}
