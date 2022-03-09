package block

type Lever struct {
	Face    string
	Facing  string
	Powered string
}

func (Lever) ID() string {
	return "minecraft:lever"
}
