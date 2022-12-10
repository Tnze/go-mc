package grids

type Smoker struct {
	*Generic
	Type int
}

func NewSmoker() *Smoker {
	return &Smoker{
		Generic: InitGenericContainer(39, 3, 1),
		Type:    21,
	}
}
