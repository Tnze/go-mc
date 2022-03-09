package block

type YellowBanner struct {
	Rotation string
}

func (YellowBanner) ID() string {
	return "minecraft:yellow_banner"
}
