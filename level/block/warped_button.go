package block

type WarpedButton struct {
	Face    string
	Facing  string
	Powered string
}

func (WarpedButton) ID() string {
	return "minecraft:warped_button"
}
