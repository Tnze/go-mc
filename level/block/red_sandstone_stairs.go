package block

type RedSandstoneStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (RedSandstoneStairs) ID() string {
	return "minecraft:red_sandstone_stairs"
}
