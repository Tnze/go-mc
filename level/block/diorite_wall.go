package block

type DioriteWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (DioriteWall) ID() string {
	return "minecraft:diorite_wall"
}
