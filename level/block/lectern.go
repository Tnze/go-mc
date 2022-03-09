package block

type Lectern struct {
	Facing   string
	Has_book string
	Powered  string
}

func (Lectern) ID() string {
	return "minecraft:lectern"
}
