package block

type AmethystCluster struct {
	Facing      string
	Waterlogged string
}

func (AmethystCluster) ID() string {
	return "minecraft:amethyst_cluster"
}
