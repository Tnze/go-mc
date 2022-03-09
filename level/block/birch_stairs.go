package block

type BirchStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (BirchStairs) ID() string {
	return "minecraft:birch_stairs"
}
