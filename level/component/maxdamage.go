package component

import pk "github.com/Tnze/go-mc/net/packet"

var _ DataComponent = (*MaxDamage)(nil)

type MaxDamage struct {
	pk.VarInt
}

// ID implements DataComponent.
func (MaxDamage) ID() string {
	return "minecraft:max_damage"
}
