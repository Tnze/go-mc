package block

type Hopper struct {
	Enabled string
	Facing  string
}

func (Hopper) ID() string {
	return "minecraft:hopper"
}
