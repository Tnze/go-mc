package block

type Rail struct {
	Shape       string
	Waterlogged string
}

func (Rail) ID() string {
	return "minecraft:rail"
}
