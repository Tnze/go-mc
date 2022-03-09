package block

type JungleDoor struct {
	Facing  string
	Half    string
	Hinge   string
	Open    string
	Powered string
}

func (JungleDoor) ID() string {
	return "minecraft:jungle_door"
}
