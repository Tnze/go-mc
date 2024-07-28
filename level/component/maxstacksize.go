package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*MaxStackSize)(nil)

type MaxStackSize struct {
	MaxStackSize pk.VarInt
}

// ID implements DataComponent.
func (MaxStackSize) ID() string {
	return "minecraft:max_stack_size"
}

// ReadFrom implements DataComponent.
func (m *MaxStackSize) ReadFrom(r io.Reader) (n int64, err error) {
	return m.MaxStackSize.ReadFrom(r)
}

// WriteTo implements DataComponent.
func (m *MaxStackSize) WriteTo(w io.Writer) (n int64, err error) {
	return m.MaxStackSize.WriteTo(w)
}
