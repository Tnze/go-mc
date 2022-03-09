package block

type Jukebox struct {
	Has_record string
}

func (Jukebox) ID() string {
	return "minecraft:jukebox"
}
