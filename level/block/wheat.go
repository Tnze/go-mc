package block

type Wheat struct {
	Age string
}

func (Wheat) ID() string {
	return "minecraft:wheat"
}
