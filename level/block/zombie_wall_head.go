package block

type ZombieWallHead struct {
	Facing string
}

func (ZombieWallHead) ID() string {
	return "minecraft:zombie_wall_head"
}
