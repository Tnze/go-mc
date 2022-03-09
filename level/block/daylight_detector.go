package block

type DaylightDetector struct {
	Inverted string
	Power    string
}

func (DaylightDetector) ID() string {
	return "minecraft:daylight_detector"
}
