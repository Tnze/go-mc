package block

type NetherBrickWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (NetherBrickWall) ID() string {
	return "minecraft:nether_brick_wall"
}
