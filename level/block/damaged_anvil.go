package block

type DamagedAnvil struct {
	Facing string
}

func (DamagedAnvil) ID() string {
	return "minecraft:damaged_anvil"
}
