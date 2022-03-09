package block

type BlackstoneStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (BlackstoneStairs) ID() string {
	return "minecraft:blackstone_stairs"
}
