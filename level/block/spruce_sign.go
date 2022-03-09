package block

type SpruceSign struct {
	Rotation    string
	Waterlogged string
}

func (SpruceSign) ID() string {
	return "minecraft:spruce_sign"
}
