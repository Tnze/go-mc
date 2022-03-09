package block

type LimeBanner struct {
	Rotation string
}

func (LimeBanner) ID() string {
	return "minecraft:lime_banner"
}
