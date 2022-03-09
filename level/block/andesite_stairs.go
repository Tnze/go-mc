package block

type AndesiteStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (AndesiteStairs) ID() string {
	return "minecraft:andesite_stairs"
}
