package block

type Lantern struct {
	Hanging     string
	Waterlogged string
}

func (Lantern) ID() string {
	return "minecraft:lantern"
}
