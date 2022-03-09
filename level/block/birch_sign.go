package block

type BirchSign struct {
	Rotation    string
	Waterlogged string
}

func (BirchSign) ID() string {
	return "minecraft:birch_sign"
}
