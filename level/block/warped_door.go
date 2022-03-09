package block

type WarpedDoor struct {
	Facing  string
	Half    string
	Hinge   string
	Open    string
	Powered string
}

func (WarpedDoor) ID() string {
	return "minecraft:warped_door"
}
