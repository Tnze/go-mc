package block

type RedBanner struct {
	Rotation string
}

func (RedBanner) ID() string {
	return "minecraft:red_banner"
}
