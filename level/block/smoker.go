package block

type Smoker struct {
	Facing string
	Lit    string
}

func (Smoker) ID() string {
	return "minecraft:smoker"
}
