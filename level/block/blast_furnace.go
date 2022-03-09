package block

type BlastFurnace struct {
	Facing string
	Lit    string
}

func (BlastFurnace) ID() string {
	return "minecraft:blast_furnace"
}
