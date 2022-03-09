package block

type GrayBanner struct {
	Rotation string
}

func (GrayBanner) ID() string {
	return "minecraft:gray_banner"
}
