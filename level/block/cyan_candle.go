package block

type CyanCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (CyanCandle) ID() string {
	return "minecraft:cyan_candle"
}
