package block

type CrimsonStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (CrimsonStairs) ID() string {
	return "minecraft:crimson_stairs"
}
