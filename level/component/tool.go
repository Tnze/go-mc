package component

import "io"

var _ DataComponent = (*Tool)(nil)

type Tool struct{}

// ID implements DataComponent.
func (Tool) ID() string {
	return "minecraft:tool"
}

// ReadFrom implements DataComponent.
func (t *Tool) ReadFrom(r io.Reader) (n int64, err error) {
	panic("unimplemented")
}

// WriteTo implements DataComponent.
func (t *Tool) WriteTo(w io.Writer) (n int64, err error) {
	panic("unimplemented")
}
