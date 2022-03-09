package block

type GrassBlock struct {
	Snowy string
}

func (GrassBlock) ID() string {
	return "minecraft:grass_block"
}
