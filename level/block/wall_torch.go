package block

type WallTorch struct {
	Facing string
}

func (WallTorch) ID() string {
	return "minecraft:wall_torch"
}
