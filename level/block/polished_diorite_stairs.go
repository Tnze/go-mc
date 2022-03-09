package block

type PolishedDioriteStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (PolishedDioriteStairs) ID() string {
	return "minecraft:polished_diorite_stairs"
}
