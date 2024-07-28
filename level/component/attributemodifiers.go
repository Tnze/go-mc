package component

import "io"

var _ DataComponent = (*AttributeModifiers)(nil)

type AttributeModifiers struct{}

// ID implements DataComponent.
func (AttributeModifiers) ID() string {
	return "minecraft:attribute_modifiers"
}

// ReadFrom implements DataComponent.
func (a *AttributeModifiers) ReadFrom(r io.Reader) (n int64, err error) {
	panic("unimplemented")
}

// WriteTo implements DataComponent.
func (a *AttributeModifiers) WriteTo(w io.Writer) (n int64, err error) {
	panic("unimplemented")
}
