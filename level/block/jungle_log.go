package block

type JungleLog struct {
	Axis string
}

func (JungleLog) ID() string {
	return "minecraft:jungle_log"
}
