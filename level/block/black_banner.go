package block

type BlackBanner struct {
	Rotation string
}

func (BlackBanner) ID() string {
	return "minecraft:black_banner"
}
