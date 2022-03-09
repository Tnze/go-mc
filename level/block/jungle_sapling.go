package block

type JungleSapling struct {
	Stage string
}

func (JungleSapling) ID() string {
	return "minecraft:jungle_sapling"
}
