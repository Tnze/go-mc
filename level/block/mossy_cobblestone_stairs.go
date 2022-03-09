package block

type MossyCobblestoneStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (MossyCobblestoneStairs) ID() string {
	return "minecraft:mossy_cobblestone_stairs"
}
