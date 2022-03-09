package block

type AcaciaButton struct {
	Face    string
	Facing  string
	Powered string
}

func (AcaciaButton) ID() string {
	return "minecraft:acacia_button"
}
