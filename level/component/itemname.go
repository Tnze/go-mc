package component

import (
	"io"

	"github.com/Tnze/go-mc/chat"
)

var _ DataComponent = (*ItemName)(nil)

type ItemName struct {
	Name chat.Message
}

// ID implements DataComponent.
func (ItemName) ID() string {
	return "minecraft:item_name"
}

// ReadFrom implements DataComponent.
func (i *ItemName) ReadFrom(r io.Reader) (n int64, err error) {
	return i.Name.ReadFrom(r)
}

// WriteTo implements DataComponent.
func (i *ItemName) WriteTo(w io.Writer) (n int64, err error) {
	return i.Name.WriteTo(w)
}
