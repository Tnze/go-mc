package block

type WarpedFenceGate struct {
	Facing  string
	In_wall string
	Open    string
	Powered string
}

func (WarpedFenceGate) ID() string {
	return "minecraft:warped_fence_gate"
}
