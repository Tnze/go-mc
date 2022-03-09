package block

type PurpleBanner struct {
	Rotation string
}

func (PurpleBanner) ID() string {
	return "minecraft:purple_banner"
}
