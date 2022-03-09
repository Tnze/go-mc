package block

type LightGrayCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (LightGrayCandle) ID() string {
	return "minecraft:light_gray_candle"
}
