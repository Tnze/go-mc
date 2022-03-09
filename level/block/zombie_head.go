package block

type ZombieHead struct {
	Rotation string
}

func (ZombieHead) ID() string {
	return "minecraft:zombie_head"
}
