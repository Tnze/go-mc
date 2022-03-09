package block

type WarpedSign struct {
	Rotation    string
	Waterlogged string
}

func (WarpedSign) ID() string {
	return "minecraft:warped_sign"
}
