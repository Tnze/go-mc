package block

type BirchTrapdoor struct {
	Facing      string
	Half        string
	Open        string
	Powered     string
	Waterlogged string
}

func (BirchTrapdoor) ID() string {
	return "minecraft:birch_trapdoor"
}
