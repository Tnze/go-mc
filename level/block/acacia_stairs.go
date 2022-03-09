package block

type AcaciaStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (AcaciaStairs) ID() string {
	return "minecraft:acacia_stairs"
}
