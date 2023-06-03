package grids

type Smoker struct {
	*Generic
}

func NewSmoker(inventory *GenericInventory) *Smoker {
	return &Smoker{InitGenericContainer("minecraft:smoker", 21, 3, inventory)}
}
