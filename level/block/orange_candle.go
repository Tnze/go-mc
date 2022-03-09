package block

type OrangeCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (OrangeCandle) ID() string {
	return "minecraft:orange_candle"
}
