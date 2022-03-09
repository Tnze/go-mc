package block

type NetherWart struct {
	Age string
}

func (NetherWart) ID() string {
	return "minecraft:nether_wart"
}
