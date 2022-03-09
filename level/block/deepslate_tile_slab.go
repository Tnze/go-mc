package block

type DeepslateTileSlab struct {
	Type        string
	Waterlogged string
}

func (DeepslateTileSlab) ID() string {
	return "minecraft:deepslate_tile_slab"
}
