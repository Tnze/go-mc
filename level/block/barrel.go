package block

type Barrel struct {
	Facing string
	Open   string
}

func (Barrel) ID() string {
	return "minecraft:barrel"
}
