package block

type Composter struct {
	Level string
}

func (Composter) ID() string {
	return "minecraft:composter"
}
