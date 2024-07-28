package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*DyedColor)(nil)

type DyedColor struct {
	RGB           pk.Int
	ShowInTooltip pk.Boolean
}

// ID implements DataComponent.
func (DyedColor) ID() string {
	return "minecraft:dyed_color"
}

// ReadFrom implements DataComponent.
func (d *DyedColor) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.Tuple{&d.RGB, &d.ShowInTooltip}.ReadFrom(r)
}

// WriteTo implements DataComponent.
func (d *DyedColor) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{&d.RGB, &d.ShowInTooltip}.WriteTo(w)
}
