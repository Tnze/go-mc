package block

type LimeCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (LimeCandle) ID() string {
	return "minecraft:lime_candle"
}
