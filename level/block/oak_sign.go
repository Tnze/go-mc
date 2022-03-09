package block

type OakSign struct {
	Rotation    string
	Waterlogged string
}

func (OakSign) ID() string {
	return "minecraft:oak_sign"
}
