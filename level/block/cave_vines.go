package block

type CaveVines struct {
	Age     string
	Berries string
}

func (CaveVines) ID() string {
	return "minecraft:cave_vines"
}
