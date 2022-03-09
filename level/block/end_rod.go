package block

type EndRod struct {
	Facing string
}

func (EndRod) ID() string {
	return "minecraft:end_rod"
}
