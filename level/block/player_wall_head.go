package block

type PlayerWallHead struct {
	Facing string
}

func (PlayerWallHead) ID() string {
	return "minecraft:player_wall_head"
}
