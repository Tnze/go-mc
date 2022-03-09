package block

type AttachedPumpkinStem struct {
	Facing string
}

func (AttachedPumpkinStem) ID() string {
	return "minecraft:attached_pumpkin_stem"
}
