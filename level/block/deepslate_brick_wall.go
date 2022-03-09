package block

type DeepslateBrickWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (DeepslateBrickWall) ID() string {
	return "minecraft:deepslate_brick_wall"
}
