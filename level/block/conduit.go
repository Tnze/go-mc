package block

type Conduit struct {
	Waterlogged string
}

func (Conduit) ID() string {
	return "minecraft:conduit"
}
