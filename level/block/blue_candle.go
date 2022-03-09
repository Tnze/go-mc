package block

type BlueCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (BlueCandle) ID() string {
	return "minecraft:blue_candle"
}
