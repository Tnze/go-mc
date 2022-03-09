package block

type SoulLantern struct {
	Hanging     string
	Waterlogged string
}

func (SoulLantern) ID() string {
	return "minecraft:soul_lantern"
}
