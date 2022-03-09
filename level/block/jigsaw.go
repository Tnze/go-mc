package block

type Jigsaw struct {
	Orientation string
}

func (Jigsaw) ID() string {
	return "minecraft:jigsaw"
}
