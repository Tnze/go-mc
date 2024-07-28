package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*MaxDamage)(nil)

type MaxDamage struct {
	MaxDamage pk.VarInt
}

// ID implements DataComponent.
func (MaxDamage) ID() string {
	return "minecraft:max_damage"
}

// ReadFrom implements DataComponent.
func (m *MaxDamage) ReadFrom(r io.Reader) (n int64, err error) {
	return m.MaxDamage.ReadFrom(r)
}

// WriteTo implements DataComponent.
func (m *MaxDamage) WriteTo(w io.Writer) (n int64, err error) {
	return m.MaxDamage.WriteTo(w)
}
