package grids

type SmithingTable struct {
	*Generic
	Type int
}

func NewSmithingTable() *SmithingTable {
	return &SmithingTable{
		Generic: InitGenericContainer(39, 3, 1),
		Type:    20,
	}
}
