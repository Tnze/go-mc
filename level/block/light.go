package block

type Light struct {
	Level       string
	Waterlogged string
}

func (Light) ID() string {
	return "minecraft:light"
}
