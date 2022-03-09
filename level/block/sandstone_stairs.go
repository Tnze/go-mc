package block

type SandstoneStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (SandstoneStairs) ID() string {
	return "minecraft:sandstone_stairs"
}
