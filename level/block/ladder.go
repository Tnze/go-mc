package block

type Ladder struct {
	Facing      string
	Waterlogged string
}

func (Ladder) ID() string {
	return "minecraft:ladder"
}
