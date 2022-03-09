package block

type DarkPrismarineStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (DarkPrismarineStairs) ID() string {
	return "minecraft:dark_prismarine_stairs"
}
