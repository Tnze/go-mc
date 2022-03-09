package block

type BrainCoralWallFan struct {
	Facing      string
	Waterlogged string
}

func (BrainCoralWallFan) ID() string {
	return "minecraft:brain_coral_wall_fan"
}
