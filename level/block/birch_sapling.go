package block

type BirchSapling struct {
	Stage string
}

func (BirchSapling) ID() string {
	return "minecraft:birch_sapling"
}
