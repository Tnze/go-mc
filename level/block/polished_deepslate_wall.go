package block

type PolishedDeepslateWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (PolishedDeepslateWall) ID() string {
	return "minecraft:polished_deepslate_wall"
}
