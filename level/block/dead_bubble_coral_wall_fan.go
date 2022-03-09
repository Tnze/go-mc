package block

type DeadBubbleCoralWallFan struct {
	Facing      string
	Waterlogged string
}

func (DeadBubbleCoralWallFan) ID() string {
	return "minecraft:dead_bubble_coral_wall_fan"
}
