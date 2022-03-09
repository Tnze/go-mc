package block

type Campfire struct {
	Facing      string
	Lit         string
	Signal_fire string
	Waterlogged string
}

func (Campfire) ID() string {
	return "minecraft:campfire"
}
