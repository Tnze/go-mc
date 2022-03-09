package block

type AcaciaDoor struct {
	Facing  string
	Half    string
	Hinge   string
	Open    string
	Powered string
}

func (AcaciaDoor) ID() string {
	return "minecraft:acacia_door"
}
