package block

type StrippedJungleLog struct {
	Axis string
}

func (StrippedJungleLog) ID() string {
	return "minecraft:stripped_jungle_log"
}
