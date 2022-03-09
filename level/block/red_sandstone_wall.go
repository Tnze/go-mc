package block

type RedSandstoneWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (RedSandstoneWall) ID() string {
	return "minecraft:red_sandstone_wall"
}
