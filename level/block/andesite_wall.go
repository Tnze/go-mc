package block

type AndesiteWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (AndesiteWall) ID() string {
	return "minecraft:andesite_wall"
}
