package block

type BlueShulkerBox struct {
	Facing string
}

func (BlueShulkerBox) ID() string {
	return "minecraft:blue_shulker_box"
}
