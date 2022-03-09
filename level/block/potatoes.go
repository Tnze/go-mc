package block

type Potatoes struct {
	Age string
}

func (Potatoes) ID() string {
	return "minecraft:potatoes"
}
