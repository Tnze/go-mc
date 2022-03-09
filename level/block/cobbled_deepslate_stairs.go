package block

type CobbledDeepslateStairs struct {
	Facing      string
	Half        string
	Shape       string
	Waterlogged string
}

func (CobbledDeepslateStairs) ID() string {
	return "minecraft:cobbled_deepslate_stairs"
}
