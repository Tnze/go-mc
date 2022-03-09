package block

type OakLog struct {
	Axis string
}

func (OakLog) ID() string {
	return "minecraft:oak_log"
}
