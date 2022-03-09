package block

type SpruceButton struct {
	Face    string
	Facing  string
	Powered string
}

func (SpruceButton) ID() string {
	return "minecraft:spruce_button"
}
