package block

type BirchFenceGate struct {
	Facing  string
	In_wall string
	Open    string
	Powered string
}

func (BirchFenceGate) ID() string {
	return "minecraft:birch_fence_gate"
}
