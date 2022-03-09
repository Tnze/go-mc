package block

type IronTrapdoor struct {
	Facing      string
	Half        string
	Open        string
	Powered     string
	Waterlogged string
}

func (IronTrapdoor) ID() string {
	return "minecraft:iron_trapdoor"
}
