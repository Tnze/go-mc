package block

type StoneButton struct {
	Face    string
	Facing  string
	Powered string
}

func (StoneButton) ID() string {
	return "minecraft:stone_button"
}
