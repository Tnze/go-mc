package block

type DeadTubeCoralWallFan struct {
	Facing      string
	Waterlogged string
}

func (DeadTubeCoralWallFan) ID() string {
	return "minecraft:dead_tube_coral_wall_fan"
}
