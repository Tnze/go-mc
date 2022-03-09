package block

type DeepslateTileStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (DeepslateTileStairs) ID() string {
	return "minecraft:deepslate_tile_stairs"
}
