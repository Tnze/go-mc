package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*RepairCost)(nil)

type RepairCost struct {
	Cost pk.VarInt
}

// ID implements DataComponent.
func (RepairCost) ID() string {
	return "minecraft:repair_cost"
}

// ReadFrom implements DataComponent.
func (r *RepairCost) ReadFrom(reader io.Reader) (n int64, err error) {
	return r.Cost.ReadFrom(reader)
}

// WriteTo implements DataComponent.
func (r *RepairCost) WriteTo(writer io.Writer) (n int64, err error) {
	return r.Cost.WriteTo(writer)
}
