package block

type WhiteBanner struct {
	Rotation string
}

func (WhiteBanner) ID() string {
	return "minecraft:white_banner"
}
