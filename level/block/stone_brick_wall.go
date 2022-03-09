package block

type StoneBrickWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (StoneBrickWall) ID() string {
	return "minecraft:stone_brick_wall"
}
