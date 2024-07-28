package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*StoredEnchantments)(nil)

type StoredEnchantments struct {
	Enchantments []struct {
		Type  pk.VarInt
		Level pk.VarInt
	}
	ShowInTooltip pk.Boolean
}

// ID implements DataComponent.
func (StoredEnchantments) ID() string {
	return "minecraft:stored_enchantments"
}

// ReadFrom implements DataComponent.
func (s *StoredEnchantments) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.Tuple{
		pk.Array(&s.Enchantments),
		&s.ShowInTooltip,
	}.ReadFrom(r)
}

// WriteTo implements DataComponent.
func (s *StoredEnchantments) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		pk.Array(&s.Enchantments),
		&s.ShowInTooltip,
	}.WriteTo(w)
}
