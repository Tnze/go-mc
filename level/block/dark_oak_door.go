package block

type DarkOakDoor struct {
	Facing  string
	Half    string
	Hinge   string
	Open    string
	Powered string
}

func (DarkOakDoor) ID() string {
	return "minecraft:dark_oak_door"
}
