package block

type JackOLantern struct {
	Facing string
}

func (JackOLantern) ID() string {
	return "minecraft:jack_o_lantern"
}
