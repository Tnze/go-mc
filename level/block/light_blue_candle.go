package block

type LightBlueCandle struct {
	Candles     string
	Lit         string
	Waterlogged string
}

func (LightBlueCandle) ID() string {
	return "minecraft:light_blue_candle"
}
