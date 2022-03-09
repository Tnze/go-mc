package block

type TrappedChest struct {
	Facing      string
	Type        string
	Waterlogged string
}

func (TrappedChest) ID() string {
	return "minecraft:trapped_chest"
}
