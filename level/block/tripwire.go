package block

type Tripwire struct {
	Attached string
	Disarmed string
	East     string
	North    string
	Powered  string
	South    string
	West     string
}

func (Tripwire) ID() string {
	return "minecraft:tripwire"
}
