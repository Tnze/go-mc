package block

type JungleSign struct {
	Rotation    string
	Waterlogged string
}

func (JungleSign) ID() string {
	return "minecraft:jungle_sign"
}
