package block

type JungleStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (JungleStairs) ID() string {
	return "minecraft:jungle_stairs"
}
