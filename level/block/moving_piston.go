package block

type MovingPiston struct {
	Facing string
	Type   string
}

func (MovingPiston) ID() string {
	return "minecraft:moving_piston"
}
