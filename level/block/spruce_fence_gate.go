package block

type SpruceFenceGate struct {
	Facing  string
	In_wall string
	Open    string
	Powered string
}

func (SpruceFenceGate) ID() string {
	return "minecraft:spruce_fence_gate"
}
