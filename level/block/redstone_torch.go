package block

type RedstoneTorch struct {
	Lit string
}

func (RedstoneTorch) ID() string {
	return "minecraft:redstone_torch"
}
