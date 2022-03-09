package block

type MushroomStem struct {
	Down  string
	East  string
	North string
	South string
	Up    string
	West  string
}

func (MushroomStem) ID() string {
	return "minecraft:mushroom_stem"
}
