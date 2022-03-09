package block

type StoneStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (StoneStairs) ID() string {
	return "minecraft:stone_stairs"
}
