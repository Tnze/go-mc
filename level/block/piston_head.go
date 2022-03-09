package block

type PistonHead struct {
	Facing string
	Short  string
	Type   string
}

func (PistonHead) ID() string {
	return "minecraft:piston_head"
}
