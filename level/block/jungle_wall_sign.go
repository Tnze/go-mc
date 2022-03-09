package block

type JungleWallSign struct {
	Facing      string
	Waterlogged string
}

func (JungleWallSign) ID() string {
	return "minecraft:jungle_wall_sign"
}
