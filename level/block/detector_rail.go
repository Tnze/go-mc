package block

type DetectorRail struct {
	Powered     string
	Shape       string
	Waterlogged string
}

func (DetectorRail) ID() string {
	return "minecraft:detector_rail"
}
