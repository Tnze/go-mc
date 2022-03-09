package block

type LimeWallBanner struct {
	Facing string
}

func (LimeWallBanner) ID() string {
	return "minecraft:lime_wall_banner"
}
