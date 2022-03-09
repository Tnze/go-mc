package block

type Beehive struct {
	Facing      string
	Honey_level string
}

func (Beehive) ID() string {
	return "minecraft:beehive"
}
