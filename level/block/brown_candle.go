package block

type BrownCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (BrownCandle) ID() string {
	return "minecraft:brown_candle"
}
