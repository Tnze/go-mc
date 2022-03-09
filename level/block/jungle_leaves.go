package block

type JungleLeaves struct {
	Distance   string
	Persistent string
}

func (JungleLeaves) ID() string {
	return "minecraft:jungle_leaves"
}
