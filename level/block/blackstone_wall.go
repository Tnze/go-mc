package block

type BlackstoneWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (BlackstoneWall) ID() string {
	return "minecraft:blackstone_wall"
}
