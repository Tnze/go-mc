package block

type ChainCommandBlock struct {
	Conditional string
	Facing      string
}

func (ChainCommandBlock) ID() string {
	return "minecraft:chain_command_block"
}
