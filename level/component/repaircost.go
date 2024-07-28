package component

import pk "github.com/Tnze/go-mc/net/packet"

var _ DataComponent = (*RepairCost)(nil)

type RepairCost struct {
	pk.VarInt
}

// ID implements DataComponent.
func (RepairCost) ID() string {
	return "minecraft:repair_cost"
}
