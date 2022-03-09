package block

type Water struct {
	Level string
}

func (Water) ID() string {
	return "minecraft:water"
}
