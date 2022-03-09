package block

type Bell struct {
	Attachment string
	Facing     string
	Powered    string
}

func (Bell) ID() string {
	return "minecraft:bell"
}
