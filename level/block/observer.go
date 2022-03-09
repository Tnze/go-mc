package block

type Observer struct {
	Facing  string
	Powered string
}

func (Observer) ID() string {
	return "minecraft:observer"
}
