package block

type Vine struct {
	East  string
	North string
	South string
	Up    string
	West  string
}

func (Vine) ID() string {
	return "minecraft:vine"
}
