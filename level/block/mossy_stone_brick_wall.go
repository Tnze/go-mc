package block

type MossyStoneBrickWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (MossyStoneBrickWall) ID() string {
	return "minecraft:mossy_stone_brick_wall"
}
