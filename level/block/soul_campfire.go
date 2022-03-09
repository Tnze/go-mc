package block

type SoulCampfire struct {
	Facing      string
	Lit         string
	Signal_fire string
	Waterlogged string
}

func (SoulCampfire) ID() string {
	return "minecraft:soul_campfire"
}
