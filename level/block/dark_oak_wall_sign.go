package block

type DarkOakWallSign struct {
	Facing      string
	Waterlogged string
}

func (DarkOakWallSign) ID() string {
	return "minecraft:dark_oak_wall_sign"
}
