package block

type WarpedWallSign struct {
	Facing      string
	Waterlogged string
}

func (WarpedWallSign) ID() string {
	return "minecraft:warped_wall_sign"
}
