package component

import "io"

var _ DataComponent = (*FireResistant)(nil)

type FireResistant struct{}

// ID implements DataComponent.
func (FireResistant) ID() string {
	return "minecraft:fire_resistant"
}

// ReadFrom implements DataComponent.
func (f *FireResistant) ReadFrom(r io.Reader) (n int64, err error) {
	return 0, nil
}

// WriteTo implements DataComponent.
func (f *FireResistant) WriteTo(w io.Writer) (n int64, err error) {
	return 0, nil
}
