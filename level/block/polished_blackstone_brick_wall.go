package block

type PolishedBlackstoneBrickWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (PolishedBlackstoneBrickWall) ID() string {
	return "minecraft:polished_blackstone_brick_wall"
}
