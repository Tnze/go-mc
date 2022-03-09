package block

type RedstoneOre struct {
	Lit string
}

func (RedstoneOre) ID() string {
	return "minecraft:redstone_ore"
}
