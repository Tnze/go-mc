package block

type Chain struct {
	Axis        string
	Waterlogged string
}

func (Chain) ID() string {
	return "minecraft:chain"
}
