package block

type CobbledDeepslateSlab struct {
	Type        string
	Waterlogged string
}

func (CobbledDeepslateSlab) ID() string {
	return "minecraft:cobbled_deepslate_slab"
}
