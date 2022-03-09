package block

type DeepslateBrickStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (DeepslateBrickStairs) ID() string {
	return "minecraft:deepslate_brick_stairs"
}
