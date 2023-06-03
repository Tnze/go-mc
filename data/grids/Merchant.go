package grids

type Merchant struct {
	*Generic
}

func NewMerchant(inventory *GenericInventory) *Merchant {
	return &Merchant{InitGenericContainer("minecraft:merchant", 18, 3, inventory)}
}
