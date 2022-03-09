package block

type EndStoneBrickStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (EndStoneBrickStairs) ID() string {
	return "minecraft:end_stone_brick_stairs"
}
