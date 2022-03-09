package block

type RedstoneWire struct {
	East  string
	North string
	Power string
	South string
	West  string
}

func (RedstoneWire) ID() string {
	return "minecraft:redstone_wire"
}
