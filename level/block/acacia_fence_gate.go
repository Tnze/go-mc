package block

type AcaciaFenceGate struct {
	Facing  string
	In_wall string
	Open    string
	Powered string
}

func (AcaciaFenceGate) ID() string {
	return "minecraft:acacia_fence_gate"
}
