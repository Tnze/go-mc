package block

type PrismarineStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (PrismarineStairs) ID() string {
	return "minecraft:prismarine_stairs"
}
