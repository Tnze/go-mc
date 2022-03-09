package block

type BirchLeaves struct {
	Distance   string
	Persistent string
}

func (BirchLeaves) ID() string {
	return "minecraft:birch_leaves"
}
