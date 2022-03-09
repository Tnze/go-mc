package block

type CobbledDeepslateWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (CobbledDeepslateWall) ID() string {
	return "minecraft:cobbled_deepslate_wall"
}
