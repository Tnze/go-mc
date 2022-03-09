package block

type LightningRod struct {
	Facing      string
	Powered     string
	Waterlogged string
}

func (LightningRod) ID() string {
	return "minecraft:lightning_rod"
}
