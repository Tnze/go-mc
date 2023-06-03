package grids

type Anvil struct {
	*Generic
}

func NewAnvil(inventory *GenericInventory) *Anvil {
	return &Anvil{InitGenericContainer("minecraft:anvil", 7, 3, inventory)}
}
