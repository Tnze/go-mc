package component

import pk "github.com/Tnze/go-mc/net/packet"

var _ DataComponent = (*MapColor)(nil)

type MapColor struct {
	// The RGB components of the color, encoded as an integer.
	pk.Int
}

// ID implements DataComponent.
func (MapColor) ID() string {
	return "minecraft:map_color"
}
