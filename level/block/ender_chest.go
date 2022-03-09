package block

type EnderChest struct {
	Facing      string
	Waterlogged string
}

func (EnderChest) ID() string {
	return "minecraft:ender_chest"
}
