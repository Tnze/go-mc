package block

type Furnace struct {
	Facing string
	Lit    string
}

func (Furnace) ID() string {
	return "minecraft:furnace"
}
