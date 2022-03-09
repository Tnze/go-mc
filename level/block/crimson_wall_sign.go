package block

type CrimsonWallSign struct {
	Facing      string
	Waterlogged string
}

func (CrimsonWallSign) ID() string {
	return "minecraft:crimson_wall_sign"
}
