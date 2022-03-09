package block

type Stonecutter struct {
	Facing string
}

func (Stonecutter) ID() string {
	return "minecraft:stonecutter"
}
