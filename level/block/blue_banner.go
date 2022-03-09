package block

type BlueBanner struct {
	Rotation string
}

func (BlueBanner) ID() string {
	return "minecraft:blue_banner"
}
