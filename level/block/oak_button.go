package block

type OakButton struct {
	Face    string
	Facing  string
	Powered string
}

func (OakButton) ID() string {
	return "minecraft:oak_button"
}
