package block

type BlackCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (BlackCandle) ID() string {
	return "minecraft:black_candle"
}
