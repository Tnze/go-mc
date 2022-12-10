package grids

type CartographyTable struct {
	*Generic
	Type int
}

func NewCartographyTable() *CartographyTable {
	return &CartographyTable{
		Generic: InitGenericContainer(39, 3, 1),
		Type:    22,
	}
}
