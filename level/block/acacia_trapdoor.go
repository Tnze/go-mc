package block

type AcaciaTrapdoor struct {
	Facing      string
	Half        string
	Open        string
	Powered     string
	Waterlogged string
}

func (AcaciaTrapdoor) ID() string {
	return "minecraft:acacia_trapdoor"
}
