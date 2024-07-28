package component

import (
	"io"

	"github.com/Tnze/go-mc/chat"
)

var _ DataComponent = (*CustomName)(nil)

type CustomName struct {
	Name chat.Message
}

// ID implements DataComponent.
func (CustomName) ID() string {
	return "minecraft:custom_name"
}

// ReadFrom implements DataComponent.
func (c *CustomName) ReadFrom(r io.Reader) (n int64, err error) {
	return c.Name.ReadFrom(r)
}

// WriteTo implements DataComponent.
func (c *CustomName) WriteTo(w io.Writer) (n int64, err error) {
	return c.Name.WriteTo(w)
}
