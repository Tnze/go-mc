package block

type DarkOakFenceGate struct {
	Facing  string
	In_wall string
	Open    string
	Powered string
}

func (DarkOakFenceGate) ID() string {
	return "minecraft:dark_oak_fence_gate"
}
