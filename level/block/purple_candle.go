package block

type PurpleCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (PurpleCandle) ID() string {
	return "minecraft:purple_candle"
}
