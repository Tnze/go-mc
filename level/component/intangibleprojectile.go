package component

import "io"

var _ DataComponent = (*IntangibleProjectile)(nil)

type IntangibleProjectile struct{}

// ID implements DataComponent.
func (IntangibleProjectile) ID() string {
	return "minecraft:intangible_projectile"
}

// ReadFrom implements DataComponent.
func (i *IntangibleProjectile) ReadFrom(r io.Reader) (n int64, err error) {
	return 0, nil
}

// WriteTo implements DataComponent.
func (i *IntangibleProjectile) WriteTo(w io.Writer) (n int64, err error) {
	return 0, nil
}
