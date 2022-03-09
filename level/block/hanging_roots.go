package block

type HangingRoots struct {
	Waterlogged string
}

func (HangingRoots) ID() string {
	return "minecraft:hanging_roots"
}
