package block

type OakWood struct {
	Axis string
}

func (OakWood) ID() string {
	return "minecraft:oak_wood"
}
