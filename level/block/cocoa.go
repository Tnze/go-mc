package block

type Cocoa struct {
	Age    string
	Facing string
}

func (Cocoa) ID() string {
	return "minecraft:cocoa"
}
