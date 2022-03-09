package block

type SandstoneWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (SandstoneWall) ID() string {
	return "minecraft:sandstone_wall"
}
