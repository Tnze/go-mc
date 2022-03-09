package block

type BirchButton struct {
	Face    string
	Facing  string
	Powered string
}

func (BirchButton) ID() string {
	return "minecraft:birch_button"
}
