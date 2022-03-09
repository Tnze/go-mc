package block

type Target struct {
	Power string
}

func (Target) ID() string {
	return "minecraft:target"
}
