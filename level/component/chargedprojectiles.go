package component

import "io"

var _ DataComponent = (*ChargedProjectiles)(nil)

type ChargedProjectiles struct {
	// TODO
}

// ID implements DataComponent.
func (ChargedProjectiles) ID() string {
	return "minecraft:charged_projectiles"
}

// ReadFrom implements DataComponent.
func (c *ChargedProjectiles) ReadFrom(r io.Reader) (n int64, err error) {
	panic("unimplemented")
}

// WriteTo implements DataComponent.
func (c *ChargedProjectiles) WriteTo(w io.Writer) (n int64, err error) {
	panic("unimplemented")
}
