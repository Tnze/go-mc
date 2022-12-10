package grids

type CraftingTable struct {
	*Generic
	Type int
}

func NewCraftingTable() *CraftingTable {
	return &CraftingTable{
		Generic: InitGenericContainer(58, 10, 1),
		Type:    11,
	}
}
