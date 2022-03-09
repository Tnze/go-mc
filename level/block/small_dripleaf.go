package block

type SmallDripleaf struct {
	Facing      string
	Half        string
	Waterlogged string
}

func (SmallDripleaf) ID() string {
	return "minecraft:small_dripleaf"
}
