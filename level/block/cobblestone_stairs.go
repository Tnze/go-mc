package block

type CobblestoneStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (CobblestoneStairs) ID() string {
	return "minecraft:cobblestone_stairs"
}
