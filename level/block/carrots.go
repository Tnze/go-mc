package block

type Carrots struct {
	Age string
}

func (Carrots) ID() string {
	return "minecraft:carrots"
}
