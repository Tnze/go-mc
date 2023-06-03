package grids

type CraftingTable struct {
	*Generic
}

func NewCraftingTable(inventory *GenericInventory) *CraftingTable {
	return &CraftingTable{InitGenericContainer("minecraft:crafting", 11, 10, inventory)}
}
