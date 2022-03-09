package block

type Dispenser struct {
	Facing    string
	Triggered string
}

func (Dispenser) ID() string {
	return "minecraft:dispenser"
}
