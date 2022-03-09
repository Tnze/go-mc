package block

type SugarCane struct {
	Age string
}

func (SugarCane) ID() string {
	return "minecraft:sugar_cane"
}
