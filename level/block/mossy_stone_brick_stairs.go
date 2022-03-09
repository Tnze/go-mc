package block

type MossyStoneBrickStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (MossyStoneBrickStairs) ID() string {
	return "minecraft:mossy_stone_brick_stairs"
}
