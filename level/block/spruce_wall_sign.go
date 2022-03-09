package block

type SpruceWallSign struct {
	Facing      string
	Waterlogged string
}

func (SpruceWallSign) ID() string {
	return "minecraft:spruce_wall_sign"
}
