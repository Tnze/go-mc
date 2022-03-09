package block

type GrayCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (GrayCandle) ID() string {
	return "minecraft:gray_candle"
}
