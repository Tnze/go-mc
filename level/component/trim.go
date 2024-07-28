package component

import "io"

var _ DataComponent = (*Trim)(nil)

type Trim struct {
}

// ID implements DataComponent.
func (Trim) ID() string {
	return "minecraft:trim"
}

// ReadFrom implements DataComponent.
func (t *Trim) ReadFrom(r io.Reader) (n int64, err error) {
	panic("unimplemented")
}

// WriteTo implements DataComponent.
func (t *Trim) WriteTo(w io.Writer) (n int64, err error) {
	panic("unimplemented")
}
