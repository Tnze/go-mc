package block

type HornCoralWallFan struct {
	Facing      string
	Waterlogged string
}

func (HornCoralWallFan) ID() string {
	return "minecraft:horn_coral_wall_fan"
}
