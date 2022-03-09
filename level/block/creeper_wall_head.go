package block

type CreeperWallHead struct {
	Facing string
}

func (CreeperWallHead) ID() string {
	return "minecraft:creeper_wall_head"
}
