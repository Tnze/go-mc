package block

type BrownMushroomBlock struct {
	Down  string
	East  string
	North string
	South string
	Up    string
	West  string
}

func (BrownMushroomBlock) ID() string {
	return "minecraft:brown_mushroom_block"
}
