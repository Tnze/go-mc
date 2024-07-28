package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*Damage)(nil)

type Damage struct {
	Damage pk.VarInt
}

// ID implements DataComponent.
func (Damage) ID() string {
	return "minecraft:damage"
}

// ReadFrom implements DataComponent.
func (d *Damage) ReadFrom(r io.Reader) (n int64, err error) {
	return d.Damage.ReadFrom(r)
}

// WriteTo implements DataComponent.
func (d *Damage) WriteTo(w io.Writer) (n int64, err error) {
	return d.Damage.WriteTo(w)
}
