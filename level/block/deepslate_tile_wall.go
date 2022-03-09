package block

type DeepslateTileWall struct {
	East        string
	North       string
	South       string
	Up          string
	Waterlogged string
	West        string
}

func (DeepslateTileWall) ID() string {
	return "minecraft:deepslate_tile_wall"
}
