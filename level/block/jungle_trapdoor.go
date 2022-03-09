package block

type JungleTrapdoor struct {
	Facing      string
	Half        string
	Open        string
	Powered     string
	Waterlogged string
}

func (JungleTrapdoor) ID() string {
	return "minecraft:jungle_trapdoor"
}
