package block

type PolishedDeepslateSlab struct {
	Type        string
	Waterlogged string
}

func (PolishedDeepslateSlab) ID() string {
	return "minecraft:polished_deepslate_slab"
}
