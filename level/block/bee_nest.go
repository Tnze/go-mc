package block

type BeeNest struct {
	Facing      string
	Honey_level string
}

func (BeeNest) ID() string {
	return "minecraft:bee_nest"
}
