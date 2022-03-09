package block

type Cake struct {
	Bites string
}

func (Cake) ID() string {
	return "minecraft:cake"
}
