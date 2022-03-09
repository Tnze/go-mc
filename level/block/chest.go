package block

type Chest struct {
	Facing      string
	Type        string
	Waterlogged string
}

func (Chest) ID() string {
	return "minecraft:chest"
}
