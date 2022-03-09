package block

type Dropper struct {
	Facing    string
	Triggered string
}

func (Dropper) ID() string {
	return "minecraft:dropper"
}
