package block

type BrownBanner struct {
	Rotation string
}

func (BrownBanner) ID() string {
	return "minecraft:brown_banner"
}
