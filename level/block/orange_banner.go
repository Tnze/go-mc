package block

type OrangeBanner struct {
	Rotation string
}

func (OrangeBanner) ID() string {
	return "minecraft:orange_banner"
}
