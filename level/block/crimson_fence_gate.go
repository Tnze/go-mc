package block

type CrimsonFenceGate struct {
	Facing  string
	In_wall string
	Open    string
	Powered string
}

func (CrimsonFenceGate) ID() string {
	return "minecraft:crimson_fence_gate"
}
