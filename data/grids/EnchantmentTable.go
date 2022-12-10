package grids

type EnchantmentTable struct {
	*Generic
	Type int
}

func NewEnchantmentTable() *EnchantmentTable {
	return &EnchantmentTable{
		Generic: InitGenericContainer(38, 2, 1),
		Type:    12,
	}
}
