package block

type OakStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (OakStairs) ID() string {
	return "minecraft:oak_stairs"
}
