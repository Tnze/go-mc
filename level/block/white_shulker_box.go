package block

type WhiteShulkerBox struct {
	Facing string
}

func (WhiteShulkerBox) ID() string {
	return "minecraft:white_shulker_box"
}
