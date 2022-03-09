package block

type SpruceDoor struct {
	Facing  string
	Half    string
	Hinge   string
	Open    string
	Powered string
}

func (SpruceDoor) ID() string {
	return "minecraft:spruce_door"
}
