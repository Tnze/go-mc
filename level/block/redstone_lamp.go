package block

type RedstoneLamp struct {
	Lit string
}

func (RedstoneLamp) ID() string {
	return "minecraft:redstone_lamp"
}
