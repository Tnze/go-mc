package block

type PolishedBlackstoneBrickStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (PolishedBlackstoneBrickStairs) ID() string {
	return "minecraft:polished_blackstone_brick_stairs"
}
