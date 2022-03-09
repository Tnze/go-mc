package block

type SpruceStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (SpruceStairs) ID() string {
	return "minecraft:spruce_stairs"
}
