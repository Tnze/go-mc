package block

type CaveVinesPlant struct {
	Berries string
}

func (CaveVinesPlant) ID() string {
	return "minecraft:cave_vines_plant"
}
