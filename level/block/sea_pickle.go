package block

type SeaPickle struct {
	Pickles     string
	Waterlogged string
}

func (SeaPickle) ID() string {
	return "minecraft:sea_pickle"
}
