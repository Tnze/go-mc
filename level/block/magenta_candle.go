package block

type MagentaCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (MagentaCandle) ID() string {
	return "minecraft:magenta_candle"
}
