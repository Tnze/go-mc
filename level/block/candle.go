package block

type Candle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (Candle) ID() string {
	return "minecraft:candle"
}
