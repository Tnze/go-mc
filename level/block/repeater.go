package block

type Repeater struct {
	Delay   string
	Facing  string
	Locked  string
	Powered string
}

func (Repeater) ID() string {
	return "minecraft:repeater"
}
