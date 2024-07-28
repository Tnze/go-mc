package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*Unbreakable)(nil)

type Unbreakable struct {
	ShowInTooltip pk.Boolean
}

// ID implements DataComponent.
func (Unbreakable) ID() string {
	return "minecraft:unbreakable"
}

// ReadFrom implements DataComponent.
func (u *Unbreakable) ReadFrom(r io.Reader) (n int64, err error) {
	return u.ShowInTooltip.ReadFrom(r)
}

// WriteTo implements DataComponent.
func (u *Unbreakable) WriteTo(w io.Writer) (n int64, err error) {
	return u.ShowInTooltip.WriteTo(w)
}
