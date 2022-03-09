package block

type AcaciaSign struct {
	Rotation    string
	Waterlogged string
}

func (AcaciaSign) ID() string {
	return "minecraft:acacia_sign"
}
