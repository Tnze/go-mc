package block

type SoulWallTorch struct {
	Facing string
}

func (SoulWallTorch) ID() string {
	return "minecraft:soul_wall_torch"
}
