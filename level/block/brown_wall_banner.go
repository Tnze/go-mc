package block

type BrownWallBanner struct {
	Facing string
}

func (BrownWallBanner) ID() string {
	return "minecraft:brown_wall_banner"
}
