package block

type DarkOakSign struct {
	Rotation    string
	Waterlogged string
}

func (DarkOakSign) ID() string {
	return "minecraft:dark_oak_sign"
}
