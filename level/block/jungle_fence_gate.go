package block

type JungleFenceGate struct {
	Facing  string
	In_wall string
	Open    string
	Powered string
}

func (JungleFenceGate) ID() string {
	return "minecraft:jungle_fence_gate"
}
