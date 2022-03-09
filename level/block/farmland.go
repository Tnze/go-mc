package block

type Farmland struct {
	Moisture string
}

func (Farmland) ID() string {
	return "minecraft:farmland"
}
