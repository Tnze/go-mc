package block

type PlayerHead struct {
	Rotation string
}

func (PlayerHead) ID() string {
	return "minecraft:player_head"
}
