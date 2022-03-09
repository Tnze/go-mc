package block

type BigDripleaf struct {
	Facing      string
	Tilt        string
	Waterlogged string
}

func (BigDripleaf) ID() string {
	return "minecraft:big_dripleaf"
}
