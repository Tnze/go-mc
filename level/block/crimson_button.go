package block

type CrimsonButton struct {
	Face    string
	Facing  string
	Powered string
}

func (CrimsonButton) ID() string {
	return "minecraft:crimson_button"
}
