package component

import "io"

var _ DataComponent = (*Enchantments)(nil)

type Enchantments struct{}

// ID implements DataComponent.
func (Enchantments) ID() string {
	return "minecraft:enchantments"
}

// ReadFrom implements DataComponent.
func (r *Enchantments) ReadFrom(reader io.Reader) (n int64, err error) {
	panic("unimplemented")
}

// WriteTo implements DataComponent.
func (r *Enchantments) WriteTo(writer io.Writer) (n int64, err error) {
	panic("unimplemented")
}
