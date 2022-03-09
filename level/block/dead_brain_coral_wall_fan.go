package block

type DeadBrainCoralWallFan struct {
	Facing      string
	Waterlogged string
}

func (DeadBrainCoralWallFan) ID() string {
	return "minecraft:dead_brain_coral_wall_fan"
}
