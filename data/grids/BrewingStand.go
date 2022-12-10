package grids

type BrewingStand struct {
	*Generic
	Type int
}

func NewBrewingStand() *BrewingStand {
	return &BrewingStand{
		Generic: InitGenericContainer(41, 5, 1),
		Type:    10,
	}
}
