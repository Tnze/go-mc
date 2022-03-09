package block

type OakDoor struct {
	Facing  string
	Half    string
	Hinge   string
	Open    string
	Powered string
}

func (OakDoor) ID() string {
	return "minecraft:oak_door"
}
