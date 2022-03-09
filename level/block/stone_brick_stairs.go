package block

type StoneBrickStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (StoneBrickStairs) ID() string {
	return "minecraft:stone_brick_stairs"
}
