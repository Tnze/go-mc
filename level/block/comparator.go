package block

type Comparator struct {
	Facing  string
	Mode    string
	Powered string
}

func (Comparator) ID() string {
	return "minecraft:comparator"
}
