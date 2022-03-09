package block

type Bamboo struct {
	Age    string
	Leaves string
	Stage  string
}

func (Bamboo) ID() string {
	return "minecraft:bamboo"
}
