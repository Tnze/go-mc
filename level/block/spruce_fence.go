package block

type SpruceFence struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (SpruceFence) ID() string {
	return "minecraft:spruce_fence"
}
