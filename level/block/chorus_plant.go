package block

type ChorusPlant struct {
	Down  string
	East  string
	North string
	South string
	Up    string
	West  string
}

func (ChorusPlant) ID() string {
	return "minecraft:chorus_plant"
}
