package block

type IronBars struct {
	East        string
	North       string
	South       string
	Waterlogged string
	West        string
}

func (IronBars) ID() string {
	return "minecraft:iron_bars"
}
