package block

type BirchWallSign struct {
	Facing      string
	Waterlogged string
}

func (BirchWallSign) ID() string {
	return "minecraft:birch_wall_sign"
}
