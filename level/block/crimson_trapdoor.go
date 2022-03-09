package block

type CrimsonTrapdoor struct {
	Facing      string
	Half        string
	Open        string
	Powered     string
	Waterlogged string
}

func (CrimsonTrapdoor) ID() string {
	return "minecraft:crimson_trapdoor"
}
