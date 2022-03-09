package block

type PoweredRail struct {
	Powered     string
	Shape       string
	Waterlogged string
}

func (PoweredRail) ID() string {
	return "minecraft:powered_rail"
}
