package grids

type EnchantmentTable struct {
	*Generic
}

func NewEnchantmentTable(inventory *GenericInventory) *EnchantmentTable {
	return &EnchantmentTable{InitGenericContainer("minecraft:enchantment", 12, 2, inventory)}
}
