package block

type IronDoor struct {
	Facing  string
	Half    string
	Hinge   string
	Open    string
	Powered string
}

func (IronDoor) ID() string {
	return "minecraft:iron_door"
}
