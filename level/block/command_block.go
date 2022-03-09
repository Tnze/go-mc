package block

type CommandBlock struct {
	Conditional string
	Facing      string
}

func (CommandBlock) ID() string {
	return "minecraft:command_block"
}
