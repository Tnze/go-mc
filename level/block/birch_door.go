package block

type BirchDoor struct {
	Facing  string
	Half    string
	Hinge   string
	Open    string
	Powered string
}

func (BirchDoor) ID() string {
	return "minecraft:birch_door"
}
