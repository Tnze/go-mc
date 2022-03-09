package block

type PolishedBlackstoneWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (PolishedBlackstoneWall) ID() string {
	return "minecraft:polished_blackstone_wall"
}
