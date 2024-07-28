package component

import "io"

var _ DataComponent = (*CanBreak)(nil)

type CanBreak struct{}

// ID implements DataComponent.
func (CanBreak) ID() string {
	return "minecraft:can_break"
}

// ReadFrom implements DataComponent.
func (c *CanBreak) ReadFrom(r io.Reader) (n int64, err error) {
	panic("unimplemented")
}

// WriteTo implements DataComponent.
func (c *CanBreak) WriteTo(w io.Writer) (n int64, err error) {
	panic("unimplemented")
}
