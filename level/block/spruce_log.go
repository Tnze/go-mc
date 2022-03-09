package block

type SpruceLog struct {
	Axis string
}

func (SpruceLog) ID() string {
	return "minecraft:spruce_log"
}
