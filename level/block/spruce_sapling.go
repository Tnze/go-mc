package block

type SpruceSapling struct {
	Stage string
}

func (SpruceSapling) ID() string {
	return "minecraft:spruce_sapling"
}
