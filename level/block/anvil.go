package block

type Anvil struct {
	Facing string
}

func (Anvil) ID() string {
	return "minecraft:anvil"
}
