package block

type WhiteCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (WhiteCandle) ID() string {
	return "minecraft:white_candle"
}
