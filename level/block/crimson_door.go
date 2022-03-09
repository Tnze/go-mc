package block

type CrimsonDoor struct {
	Facing  string
	Half    string
	Hinge   string
	Open    string
	Powered string
}

func (CrimsonDoor) ID() string {
	return "minecraft:crimson_door"
}
