package block

type EndStoneBrickWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (EndStoneBrickWall) ID() string {
	return "minecraft:end_stone_brick_wall"
}
