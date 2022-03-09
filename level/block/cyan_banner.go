package block

type CyanBanner struct {
	Rotation string
}

func (CyanBanner) ID() string {
	return "minecraft:cyan_banner"
}
