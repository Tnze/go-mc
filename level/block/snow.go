package block

type Snow struct {
	Layers string
}

func (Snow) ID() string {
	return "minecraft:snow"
}
