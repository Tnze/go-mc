package block

type CrimsonSign struct {
	Rotation    string
	Waterlogged string
}

func (CrimsonSign) ID() string {
	return "minecraft:crimson_sign"
}
