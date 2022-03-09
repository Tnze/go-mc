package block

type GraniteWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (GraniteWall) ID() string {
	return "minecraft:granite_wall"
}
