package block

type Scaffolding struct {
	Bottom      string
	Distance    string
	Waterlogged string
}

func (Scaffolding) ID() string {
	return "minecraft:scaffolding"
}
