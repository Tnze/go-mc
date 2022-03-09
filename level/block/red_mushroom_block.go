package block

type RedMushroomBlock struct {
	Down  string
	East  string
	North string
	South string
	Up    string
	West  string
}

func (RedMushroomBlock) ID() string {
	return "minecraft:red_mushroom_block"
}
