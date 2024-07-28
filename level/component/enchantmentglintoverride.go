package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*EnchantmentGlintOverride)(nil)

type EnchantmentGlintOverride struct {
	HasGlint pk.Boolean
}

// ID implements DataComponent.
func (EnchantmentGlintOverride) ID() string {
	return "minecraft:enchantment_glint_override"
}

// ReadFrom implements DataComponent.
func (e *EnchantmentGlintOverride) ReadFrom(r io.Reader) (n int64, err error) {
	return e.HasGlint.ReadFrom(r)
}

// WriteTo implements DataComponent.
func (e *EnchantmentGlintOverride) WriteTo(w io.Writer) (n int64, err error) {
	return e.HasGlint.WriteTo(w)
}
