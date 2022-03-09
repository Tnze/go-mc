package block

type PolishedBlackstoneButton struct {
	Face    string
	Facing  string
	Powered string
}

func (PolishedBlackstoneButton) ID() string {
	return "minecraft:polished_blackstone_button"
}
