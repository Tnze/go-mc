package block

type PinkWallBanner struct {
	Facing string
}

func (PinkWallBanner) ID() string {
	return "minecraft:pink_wall_banner"
}
