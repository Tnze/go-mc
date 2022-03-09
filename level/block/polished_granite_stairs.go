package block

type PolishedGraniteStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (PolishedGraniteStairs) ID() string {
	return "minecraft:polished_granite_stairs"
}
