package block

type GreenBanner struct {
	Rotation string
}

func (GreenBanner) ID() string {
	return "minecraft:green_banner"
}
