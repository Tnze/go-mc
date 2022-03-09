package block

type DeadFireCoralWallFan struct {
	Facing      string
	Waterlogged string
}

func (DeadFireCoralWallFan) ID() string {
	return "minecraft:dead_fire_coral_wall_fan"
}
