package block

type CyanWallBanner struct {
	Facing string
}

func (CyanWallBanner) ID() string {
	return "minecraft:cyan_wall_banner"
}
