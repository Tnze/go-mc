package block

type SpruceTrapdoor struct {
	Facing      string
	Half        string
	Open        string
	Powered     string
	Waterlogged string
}

func (SpruceTrapdoor) ID() string {
	return "minecraft:spruce_trapdoor"
}
