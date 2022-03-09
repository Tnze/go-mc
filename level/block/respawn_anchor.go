package block

type RespawnAnchor struct {
	Charges string
}

func (RespawnAnchor) ID() string {
	return "minecraft:respawn_anchor"
}
