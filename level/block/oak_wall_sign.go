package block

type OakWallSign struct {
	Facing      string
	Waterlogged string
}

func (OakWallSign) ID() string {
	return "minecraft:oak_wall_sign"
}
