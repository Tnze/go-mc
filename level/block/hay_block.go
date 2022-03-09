package block

type HayBlock struct {
	Axis string
}

func (HayBlock) ID() string {
	return "minecraft:hay_block"
}
