package block

type Lava struct {
	Level string
}

func (Lava) ID() string {
	return "minecraft:lava"
}
