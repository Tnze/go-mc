package block

type TurtleEgg struct {
	Eggs  string
	Hatch string
}

func (TurtleEgg) ID() string {
	return "minecraft:turtle_egg"
}
