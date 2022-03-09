package block

type NetherPortal struct {
	Axis string
}

func (NetherPortal) ID() string {
	return "minecraft:nether_portal"
}
