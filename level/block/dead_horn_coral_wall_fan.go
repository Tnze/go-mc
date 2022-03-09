package block

type DeadHornCoralWallFan struct {
	Facing      string
	Waterlogged string
}

func (DeadHornCoralWallFan) ID() string {
	return "minecraft:dead_horn_coral_wall_fan"
}
