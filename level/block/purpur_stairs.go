package block

type PurpurStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (PurpurStairs) ID() string {
	return "minecraft:purpur_stairs"
}
