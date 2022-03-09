package block

type TripwireHook struct {
	Attached string
	Facing   string
	Powered  string
}

func (TripwireHook) ID() string {
	return "minecraft:tripwire_hook"
}
