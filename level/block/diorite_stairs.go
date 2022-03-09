package block

type DioriteStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (DioriteStairs) ID() string {
	return "minecraft:diorite_stairs"
}
