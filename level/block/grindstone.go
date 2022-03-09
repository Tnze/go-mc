package block

type Grindstone struct {
	Face   string
	Facing string
}

func (Grindstone) ID() string {
	return "minecraft:grindstone"
}
