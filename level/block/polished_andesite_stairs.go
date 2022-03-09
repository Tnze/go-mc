package block

type PolishedAndesiteStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (PolishedAndesiteStairs) ID() string {
	return "minecraft:polished_andesite_stairs"
}
