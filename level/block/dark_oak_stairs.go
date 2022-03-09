package block

type DarkOakStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (DarkOakStairs) ID() string {
	return "minecraft:dark_oak_stairs"
}
