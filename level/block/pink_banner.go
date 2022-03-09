package block

type PinkBanner struct {
	Rotation string
}

func (PinkBanner) ID() string {
	return "minecraft:pink_banner"
}
