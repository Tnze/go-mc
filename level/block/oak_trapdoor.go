package block

type OakTrapdoor struct {
	Facing      string
	Half        string
	Open        string
	Powered     string
	Waterlogged string
}

func (OakTrapdoor) ID() string {
	return "minecraft:oak_trapdoor"
}
