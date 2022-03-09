package block

type RedstoneWallTorch struct {
	Facing string
	Lit    string
}

func (RedstoneWallTorch) ID() string {
	return "minecraft:redstone_wall_torch"
}
