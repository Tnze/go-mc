package block

type FireCoralWallFan struct {
	Facing      string
	Waterlogged string
}

func (FireCoralWallFan) ID() string {
	return "minecraft:fire_coral_wall_fan"
}
