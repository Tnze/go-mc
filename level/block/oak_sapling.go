package block

type OakSapling struct {
	Stage string
}

func (OakSapling) ID() string {
	return "minecraft:oak_sapling"
}
