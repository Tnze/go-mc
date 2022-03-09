package block

type DragonWallHead struct {
	Facing string
}

func (DragonWallHead) ID() string {
	return "minecraft:dragon_wall_head"
}
