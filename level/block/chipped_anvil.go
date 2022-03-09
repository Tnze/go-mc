package block

type ChippedAnvil struct {
	Facing string
}

func (ChippedAnvil) ID() string {
	return "minecraft:chipped_anvil"
}
