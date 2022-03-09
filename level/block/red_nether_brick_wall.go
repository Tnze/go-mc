package block

type RedNetherBrickWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (RedNetherBrickWall) ID() string {
	return "minecraft:red_nether_brick_wall"
}
