package block

type DarkOakTrapdoor struct {
	Facing      string
	Half        string
	Open        string
	Powered     string
	Waterlogged string
}

func (DarkOakTrapdoor) ID() string {
	return "minecraft:dark_oak_trapdoor"
}
