package block

type StickyPiston struct {
	Extended string
	Facing   string
}

func (StickyPiston) ID() string {
	return "minecraft:sticky_piston"
}
