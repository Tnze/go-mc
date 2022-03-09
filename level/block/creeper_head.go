package block

type CreeperHead struct {
	Rotation string
}

func (CreeperHead) ID() string {
	return "minecraft:creeper_head"
}
