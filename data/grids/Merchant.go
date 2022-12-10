package grids

type Merchant struct {
	*Generic
	Type int
}

func NewMerchant() *Merchant {
	return &Merchant{
		Generic: InitGenericContainer(39, 3, 1),
		Type:    18,
	}
}
