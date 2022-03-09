package block

type GrayShulkerBox struct {
	Facing string
}

func (GrayShulkerBox) ID() string {
	return "minecraft:gray_shulker_box"
}
