package component

import (
	"io"

	"github.com/Tnze/go-mc/nbt/dynbt"
	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*EntityData)(nil)

type EntityData struct {
	dynbt.Value
}

// ID implements DataComponent.
func (EntityData) ID() string {
	return "minecraft:entity_data"
}

// ReadFrom implements DataComponent.
func (e *EntityData) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.NBT(&e.Value).ReadFrom(r)
}

// WriteTo implements DataComponent.
func (e *EntityData) WriteTo(w io.Writer) (n int64, err error) {
	return pk.NBT(&e.Value).WriteTo(w)
}
