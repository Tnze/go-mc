package component

import (
	"io"

	"github.com/Tnze/go-mc/nbt/dynbt"
	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*MapDecorations)(nil)

type MapDecorations struct {
	dynbt.Value
}

// ID implements DataComponent.
func (MapDecorations) ID() string {
	return "minecraft:map_decorations"
}

// ReadFrom implements DataComponent.
func (m *MapDecorations) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.NBT(&m.Value).ReadFrom(r)
}

// WriteTo implements DataComponent.
func (m *MapDecorations) WriteTo(w io.Writer) (n int64, err error) {
	return pk.NBT(&m.Value).WriteTo(w)
}
