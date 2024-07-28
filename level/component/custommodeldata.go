package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*CustomModelData)(nil)

type CustomModelData struct {
	Value pk.VarInt
}

// ID implements DataComponent.
func (CustomModelData) ID() string {
	return "minecraft:custom_model_data"
}

// ReadFrom implements DataComponent.
func (c *CustomModelData) ReadFrom(r io.Reader) (n int64, err error) {
	return c.Value.ReadFrom(r)
}

// WriteTo implements DataComponent.
func (c *CustomModelData) WriteTo(w io.Writer) (n int64, err error) {
	return c.Value.WriteTo(w)
}
