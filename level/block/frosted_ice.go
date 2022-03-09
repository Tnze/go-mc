package block

type FrostedIce struct {
	Age string
}

func (FrostedIce) ID() string {
	return "minecraft:frosted_ice"
}
