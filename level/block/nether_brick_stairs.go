package block

type NetherBrickStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (NetherBrickStairs) ID() string {
	return "minecraft:nether_brick_stairs"
}
