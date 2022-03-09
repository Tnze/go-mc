package block

type StructureBlock struct {
	Mode string
}

func (StructureBlock) ID() string {
	return "minecraft:structure_block"
}
