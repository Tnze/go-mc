package block

type RedNetherBrickStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (RedNetherBrickStairs) ID() string {
	return "minecraft:red_nether_brick_stairs"
}
