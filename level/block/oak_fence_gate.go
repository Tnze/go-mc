package block

type OakFenceGate struct {
	Facing  string
	In_wall string
	Open    string
	Powered string
}

func (OakFenceGate) ID() string {
	return "minecraft:oak_fence_gate"
}
