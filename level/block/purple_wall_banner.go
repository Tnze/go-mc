package block

type PurpleWallBanner struct {
	Facing string
}

func (PurpleWallBanner) ID() string {
	return "minecraft:purple_wall_banner"
}
