package block

type DarkOakButton struct {
	Face    string
	Facing  string
	Powered string
}

func (DarkOakButton) ID() string {
	return "minecraft:dark_oak_button"
}
