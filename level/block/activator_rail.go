package block

type ActivatorRail struct {
	Powered     string
	Shape       string
	Waterlogged string
}

func (ActivatorRail) ID() string {
	return "minecraft:activator_rail"
}
