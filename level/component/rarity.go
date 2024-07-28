package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*Rarity)(nil)

type Rarity int32

const (
	Common Rarity = iota
	Uncommon
	Rare
	Epic
)

// ID implements DataComponent.
func (Rarity) ID() string {
	return "minecraft:rarity"
}

// ReadFrom implements DataComponent.
func (r *Rarity) ReadFrom(reader io.Reader) (n int64, err error) {
	return (*pk.VarInt)(r).ReadFrom(reader)
}

// WriteTo implements DataComponent.
func (r *Rarity) WriteTo(writer io.Writer) (n int64, err error) {
	return (*pk.VarInt)(r).WriteTo(writer)
}
