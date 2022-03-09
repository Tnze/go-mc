package block

type BrickWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (BrickWall) ID() string {
	return "minecraft:brick_wall"
}
