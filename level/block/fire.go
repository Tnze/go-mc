package block

type Fire struct {
	Age   string
	East  string
	North string
	South string
	Up    string
	West  string
}

func (Fire) ID() string {
	return "minecraft:fire"
}
