package block

type PrismarineBrickStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (PrismarineBrickStairs) ID() string {
	return "minecraft:prismarine_brick_stairs"
}
