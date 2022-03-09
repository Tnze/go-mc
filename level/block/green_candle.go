package block

type GreenCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (GreenCandle) ID() string {
	return "minecraft:green_candle"
}
