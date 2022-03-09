package block

type RepeatingCommandBlock struct {
	Conditional string
	Facing      string
}

func (RepeatingCommandBlock) ID() string {
	return "minecraft:repeating_command_block"
}
