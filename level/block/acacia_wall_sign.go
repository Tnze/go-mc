package block

type AcaciaWallSign struct {
	Facing      string
	Waterlogged string
}

func (AcaciaWallSign) ID() string {
	return "minecraft:acacia_wall_sign"
}
