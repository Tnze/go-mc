package block

type RedCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (RedCandle) ID() string {
	return "minecraft:red_candle"
}
