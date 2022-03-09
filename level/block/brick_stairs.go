package block

type BrickStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (BrickStairs) ID() string {
	return "minecraft:brick_stairs"
}
