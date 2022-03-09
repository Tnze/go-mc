package block

type JungleButton struct {
	Face    string
	Facing  string
	Powered string
}

func (JungleButton) ID() string {
	return "minecraft:jungle_button"
}
