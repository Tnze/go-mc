package block

type ShulkerBox struct {
	Facing string
}

func (ShulkerBox) ID() string {
	return "minecraft:shulker_box"
}
