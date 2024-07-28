package component

import (
	"io"

	"github.com/Tnze/go-mc/chat"
	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*Lore)(nil)

type Lore struct {
	Lines []chat.Message
}

// ID implements DataComponent.
func (Lore) ID() string {
	return "minecraft:lore"
}

// ReadFrom implements DataComponent.
func (l *Lore) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.Array(&l.Lines).ReadFrom(r)
}

// WriteTo implements DataComponent.
func (l *Lore) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Array(&l.Lines).WriteTo(w)
}
