package block

type MossyCobblestoneWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (MossyCobblestoneWall) ID() string {
	return "minecraft:mossy_cobblestone_wall"
}
