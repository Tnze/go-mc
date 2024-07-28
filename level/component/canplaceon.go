package component

import "io"

var _ DataComponent = (*CanPlaceOn)(nil)

type CanPlaceOn struct{}

// ID implements DataComponent.
func (CanPlaceOn) ID() string {
	return "minecraft:can_place_on"
}

// ReadFrom implements DataComponent.
func (c *CanPlaceOn) ReadFrom(r io.Reader) (n int64, err error) {
	panic("unimplemented")
}

// WriteTo implements DataComponent.
func (c *CanPlaceOn) WriteTo(w io.Writer) (n int64, err error) {
	panic("unimplemented")
}
