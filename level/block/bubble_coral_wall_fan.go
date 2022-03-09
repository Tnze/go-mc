package block

type BubbleCoralWallFan struct {
	Facing      string
	Waterlogged string
}

func (BubbleCoralWallFan) ID() string {
	return "minecraft:bubble_coral_wall_fan"
}
