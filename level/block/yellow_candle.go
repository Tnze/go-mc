package block

type YellowCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (YellowCandle) ID() string {
	return "minecraft:yellow_candle"
}
