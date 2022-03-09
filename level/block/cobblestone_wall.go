package block

type CobblestoneWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (CobblestoneWall) ID() string {
	return "minecraft:cobblestone_wall"
}
