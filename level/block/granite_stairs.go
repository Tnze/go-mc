package block

type GraniteStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (GraniteStairs) ID() string {
	return "minecraft:granite_stairs"
}
