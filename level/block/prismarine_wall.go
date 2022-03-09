package block

type PrismarineWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (PrismarineWall) ID() string {
	return "minecraft:prismarine_wall"
}
