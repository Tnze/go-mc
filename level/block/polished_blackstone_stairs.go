package block

type PolishedBlackstoneStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (PolishedBlackstoneStairs) ID() string {
	return "minecraft:polished_blackstone_stairs"
}
