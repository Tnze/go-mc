package block

type Piston struct {
	Extended string
	Facing   string
}

func (Piston) ID() string {
	return "minecraft:piston"
}
