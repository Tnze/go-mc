package block

type PinkCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (PinkCandle) ID() string {
	return "minecraft:pink_candle"
}
