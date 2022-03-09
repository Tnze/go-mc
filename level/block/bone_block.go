package block

type BoneBlock struct {
	Axis string
}

func (BoneBlock) ID() string {
	return "minecraft:bone_block"
}
