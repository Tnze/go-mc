package block

type WarpedTrapdoor struct {
	Facing      string
	Half        string
	Open        string
	Powered     string
	Waterlogged string
}

func (WarpedTrapdoor) ID() string {
	return "minecraft:warped_trapdoor"
}
