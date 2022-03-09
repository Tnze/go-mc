package block

type TubeCoralWallFan struct {
	Facing      string
	Waterlogged string
}

func (TubeCoralWallFan) ID() string {
	return "minecraft:tube_coral_wall_fan"
}
